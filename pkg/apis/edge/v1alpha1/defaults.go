/* SPDX-License-Identifier: LGPL-2.1-or-later */

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

// SetDefaults_EdgeDeviceSpec sets defaults for EdgeDevice spec
func SetDefaults_EdgeDeviceSpec(obj *EdgeDeviceSpec) {

}
