/* SPDX-License-Identifier: LGPL-2.1-or-later */

package edge

import (
	"github.com/engelmi/edge-api-server/pkg/apis/edge"
	"github.com/engelmi/edge-api-server/pkg/apis/edge/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

// Install registers the API group and adds types to a scheme
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(edge.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(scheme.SetVersionPriority(v1alpha1.SchemeGroupVersion))
}
