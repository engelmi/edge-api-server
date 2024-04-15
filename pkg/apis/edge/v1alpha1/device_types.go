/* SPDX-License-Identifier: LGPL-2.1-or-later */

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type EdgeDeviceStatus string

const (
	SystemStatusUp       = EdgeDeviceStatus("up")
	SystemStatusDegraded = EdgeDeviceStatus("degraded")
	SystemStatusDown     = EdgeDeviceStatus("down")
)

type NodeStatus string

const (
	NodeStatusOnline  = NodeStatus("online")
	NodeStatusOffline = NodeStatus("offline")
)

type EdgeNode struct {
	Name              string     `json:"name,omitempty"`
	Status            NodeStatus `json:"status,omitempty"`
	LastSeenTimestamp string     `json:"lastSeenTimestamp,omitempty"`
}

type EdgeNodes []EdgeNode

type EdgeDeviceSpec struct {
	Nodes EdgeNodes `json:"nodes,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EdgeDevice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeDeviceSpec   `json:"spec,omitempty"`
	Status EdgeDeviceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EdgeDeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []EdgeDevice `json:"items"`
}
