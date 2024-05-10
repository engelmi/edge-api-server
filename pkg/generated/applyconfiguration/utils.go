/* SPDX-License-Identifier: LGPL-2.1-or-later */

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1alpha1 "github.com/engelmi/edge-api-server/pkg/apis/edge/v1alpha1"
	edgev1alpha1 "github.com/engelmi/edge-api-server/pkg/generated/applyconfiguration/edge/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=edge, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("EdgeDevice"):
		return &edgev1alpha1.EdgeDeviceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EdgeDeviceSpec"):
		return &edgev1alpha1.EdgeDeviceSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EdgeNode"):
		return &edgev1alpha1.EdgeNodeApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Workload"):
		return &edgev1alpha1.WorkloadApplyConfiguration{}

	}
	return nil
}
