/* SPDX-License-Identifier: LGPL-2.1-or-later */

package device

import (
	"github.com/engelmi/edge-api-server/pkg/apis/edge"
	"github.com/engelmi/edge-api-server/pkg/registry"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

// NewREST returns a RESTStorage object that will work against API services.
func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)

	store := &genericregistry.Store{
		NewFunc:     func() runtime.Object { return &edge.EdgeDevice{} },
		NewListFunc: func() runtime.Object { return &edge.EdgeDeviceList{} },

		PredicateFunc:            MatchEdgeDevice,
		DefaultQualifiedResource: edge.Resource("edgedevices"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		TableConvertor: rest.NewDefaultTableConvertor(edge.Resource("edgedevices")),
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &registry.REST{Store: store}, nil
}
