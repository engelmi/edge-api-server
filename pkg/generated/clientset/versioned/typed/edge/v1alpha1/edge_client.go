/* SPDX-License-Identifier: LGPL-2.1-or-later */

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/engelmi/edge-api-server/pkg/apis/edge/v1alpha1"
	"github.com/engelmi/edge-api-server/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type EdgeV1alpha1Interface interface {
	RESTClient() rest.Interface
	EdgeDevicesGetter
}

// EdgeV1alpha1Client is used to interact with features provided by the edge group.
type EdgeV1alpha1Client struct {
	restClient rest.Interface
}

func (c *EdgeV1alpha1Client) EdgeDevices(namespace string) EdgeDeviceInterface {
	return newEdgeDevices(c, namespace)
}

// NewForConfig creates a new EdgeV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*EdgeV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new EdgeV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*EdgeV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &EdgeV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new EdgeV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *EdgeV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new EdgeV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *EdgeV1alpha1Client {
	return &EdgeV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *EdgeV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
