/* SPDX-License-Identifier: LGPL-2.1-or-later */

package device

import (
	"context"
	"fmt"

	"github.com/engelmi/edge-api-server/pkg/apis/edge"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
)

// NewStrategy creates and returns a deviceStrategy instance
func NewStrategy(typer runtime.ObjectTyper) deviceStrategy {
	return deviceStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not an EdgeDevice
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*edge.EdgeDevice)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not an EdgeDevice")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

// MatchEdgeDevice is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchEdgeDevice(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *edge.EdgeDevice) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type deviceStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (deviceStrategy) NamespaceScoped() bool {
	return true
}

func (deviceStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (deviceStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (deviceStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (deviceStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (deviceStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (deviceStrategy) Canonicalize(obj runtime.Object) {
}

func (deviceStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// WarningsOnCreate returns warnings for the creation of the given object.
func (deviceStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

// WarningsOnUpdate returns warnings for the given update.
func (deviceStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}
