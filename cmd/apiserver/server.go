/* SPDX-License-Identifier: LGPL-2.1-or-later */

package main

import (
	"fmt"
	"io"
	"net"

	"github.com/engelmi/edge-api-server/pkg/apis/edge"
	"github.com/spf13/cobra"

	"k8s.io/api/node/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
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
			Codecs.LegacyCodec(edge.SchemeGroupVersion),
		),

		StdOut: out,
		StdErr: errOut,
	}
	o.RecommendedOptions.Etcd.StorageConfig.EncodeVersioner = runtime.NewMultiGroupVersioner(v1alpha1.SchemeGroupVersion, schema.GroupKind{Group: v1alpha1.GroupName})
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
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", o.AlternateDNS, []net.IP{netutils.ParseIPSloppy("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	serverConfig := genericapiserver.NewRecommendedConfig(Codecs)
	if err := o.RecommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	config := &Config{
		GenericConfig: serverConfig,
		ExtraConfig:   ExtraConfig{},
	}
	return config, nil
}

// RunEdgeServer starts a new EdgeServer given EdgeServerOptions
func (o EdgeServerOptions) RunEdgeServer(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New()
	if err != nil {
		return err
	}

	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}
