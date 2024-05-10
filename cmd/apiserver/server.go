/* SPDX-License-Identifier: LGPL-2.1-or-later */

package main

import (
	"fmt"
	"io"
	"net"

	"github.com/spf13/cobra"

	edgeapi "github.com/engelmi/edge-api-server/pkg/apis/edge"
	edgeapi_v1alpha1 "github.com/engelmi/edge-api-server/pkg/apis/edge/v1alpha1"
	"github.com/engelmi/edge-api-server/pkg/bridges"
	"github.com/engelmi/edge-api-server/pkg/generated/clientset/versioned/typed/edge/v1alpha1"
	edge_openapi "github.com/engelmi/edge-api-server/pkg/generated/openapi"
	"github.com/engelmi/edge-api-server/pkg/registry"
	"github.com/engelmi/edge-api-server/pkg/registry/edge/device"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apiserver/pkg/endpoints/openapi"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	k8srest "k8s.io/client-go/rest"
	k8sclientcmd "k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	netutils "k8s.io/utils/net"
)

const defaultEtcdPathPrefix = "/registry/org.redhat.edgeapi"

// EdgeServerOptions contains state for master/api server
type EdgeServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions

	StdOut io.Writer
	StdErr io.Writer

	AlternateDNS []string
}

// NewEdgeServerOptions returns a new EdgeServerOptions
func NewEdgeServerOptions(out, errOut io.Writer) *EdgeServerOptions {
	o := &EdgeServerOptions{
		RecommendedOptions: genericoptions.NewRecommendedOptions(
			defaultEtcdPathPrefix,
			Codecs.LegacyCodec(edgeapi_v1alpha1.SchemeGroupVersion),
		),

		StdOut: out,
		StdErr: errOut,
	}
	o.RecommendedOptions.Etcd.StorageConfig.EncodeVersioner = runtime.NewMultiGroupVersioner(
		edgeapi_v1alpha1.SchemeGroupVersion,
		schema.GroupKind{Group: edgeapi_v1alpha1.GroupName},
	)

	return o
}

// NewCommandStartEdgeServer provides a CLI handler for 'start master' command
// with a default EdgeServerOptions.
func NewCommandStartEdgeServer(defaults *EdgeServerOptions, stopCh <-chan struct{}) *cobra.Command {
	o := *defaults
	cmd := &cobra.Command{
		Short: "Launch an Edge API server",
		Long:  "Launch an Edge API server",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(args); err != nil {
				return err
			}

			config, err := k8sclientcmd.BuildConfigFromFlags("", "")
			if err != nil {
				klog.Errorf("Failed to create k8s config: %s", err)
			}
			mqttBridge := bridges.NewMQTTBridge(
				c.Context(),
				"edgeapi",
				bridges.NewMQTTClient(
					"mosquitto-mqtts.edgeapi-mqtt.svc.cluster.local",
					8883,
					"someuser",
					"topsecret",
				),
				v1alpha1.NewForConfigOrDie(
					k8srest.AddUserAgent(config, ""),
				),
			)
			go mqttBridge.Run(stopCh)

			if err := o.RunEdgeServer(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	o.RecommendedOptions.AddFlags(flags)
	utilfeature.DefaultMutableFeatureGate.AddFlag(flags)

	return cmd
}

// Validate validates EdgeServerOptions
func (o EdgeServerOptions) Validate(args []string) error {
	errors := []error{}
	errors = append(errors, o.RecommendedOptions.Validate()...)
	return utilerrors.NewAggregate(errors)
}

// Complete fills in fields required to have valid data
func (o *EdgeServerOptions) Complete() error {
	return nil
}

// Config returns config for the api server given EdgeServerOptions
func (o *EdgeServerOptions) Config() (*Config, error) {
	// TODO: have a "real" external address
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts(
		"localhost",
		o.AlternateDNS,
		[]net.IP{netutils.ParseIPSloppy("127.0.0.1")},
	); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	serverConfig := genericapiserver.NewRecommendedConfig(Codecs)
	if err := o.RecommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	serverConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(
		edge_openapi.GetOpenAPIDefinitions,
		openapi.NewDefinitionNamer(Scheme),
	)
	serverConfig.OpenAPIConfig.Info.Title = "EdgeAPI"
	serverConfig.OpenAPIConfig.Info.Version = "0.1"

	serverConfig.OpenAPIV3Config = genericapiserver.DefaultOpenAPIV3Config(
		edge_openapi.GetOpenAPIDefinitions,
		openapi.NewDefinitionNamer(Scheme),
	)
	serverConfig.OpenAPIV3Config.Info.Title = "EdgeAPI"
	serverConfig.OpenAPIV3Config.Info.Version = "0.1"

	config := &Config{
		GenericConfig: serverConfig,
		ExtraConfig:   ExtraConfig{},
	}
	return config, nil
}

// EdgeServer contains state for a Kubernetes cluster master/api server.
type EdgeServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

// New returns a new instance of EdgeServer from the given config.
func NewEdgeServer(c *CompletedConfig) (*EdgeServer, error) {
	genericServer, err := c.GenericConfig.New("edge-api-server", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	s := &EdgeServer{
		GenericAPIServer: genericServer,
	}

	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(edgeapi.GroupName, Scheme, metav1.ParameterCodec, Codecs)

	v1alpha1storage := map[string]rest.Storage{}
	v1alpha1storage["edgedevices"] = registry.RESTInPeace(device.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter))
	apiGroupInfo.VersionedResourcesStorageMap["v1alpha1"] = v1alpha1storage

	if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}

	return s, nil
}

// RunEdgeServer starts a new EdgeServer given EdgeServerOptions
func (o EdgeServerOptions) RunEdgeServer(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := NewEdgeServer(config.Complete())
	if err != nil {
		return err
	}

	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}
