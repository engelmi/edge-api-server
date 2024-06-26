/* SPDX-License-Identifier: LGPL-2.1-or-later */

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type EdgeDeviceStatus string

const (
	EdgeDeviceStatusUp       = EdgeDeviceStatus("up")
	EdgeDeviceStatusDegraded = EdgeDeviceStatus("degraded")
	EdgeDeviceStatusDown     = EdgeDeviceStatus("down")
)

type Workload struct {
	Name     string `json:"name,omitempty"`
	State    string `json:"state,omitempty"`
	SubState string `json:"substate,omitempty"`
}

type EdgeNodeStatus string

const (
	EdgeNodeStatusOnline  = EdgeNodeStatus("online")
	EdgeNodeStatusOffline = EdgeNodeStatus("offline")
)

type EdgeNode struct {
	Name              string         `json:"name,omitempty"`
	Status            EdgeNodeStatus `json:"status,omitempty"`
	LastSeenTimestamp string         `json:"lastSeenTimestamp,omitempty"`

	Workloads []Workload `json:"workloads,omitempty"`
}

type EdgeNodes []EdgeNode

type EdgeDeviceSpec struct {
	ID    string    `json:"id,omitempty"`
	Type  string    `json:"type,omitempty"`
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
