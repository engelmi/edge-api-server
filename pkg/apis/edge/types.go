/* SPDX-License-Identifier: LGPL-2.1-or-later */

package edge

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type EdgeDeviceStatus string

const (
	EdgeDeviceStatusUp       = EdgeDeviceStatus("up")
	EdgeDeviceStatusDegraded = EdgeDeviceStatus("degraded")
	EdgeDeviceStatusDown     = EdgeDeviceStatus("down")
)

type EdgeNodeStatus string

const (
	EdgeNodeStatusOnline  = EdgeNodeStatus("online")
	EdgeNodeStatusOffline = EdgeNodeStatus("offline")
)

type EdgeNode struct {
	Name              string
	Status            EdgeNodeStatus
	LastSeenTimestamp string
}

type EdgeNodes []EdgeNode

type EdgeDeviceSpec struct {
	Nodes EdgeNodes
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EdgeDevice struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   EdgeDeviceSpec
	Status EdgeDeviceStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EdgeDeviceList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []EdgeDevice
}
