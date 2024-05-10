//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/* SPDX-License-Identifier: LGPL-2.1-or-later */

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	edge "github.com/engelmi/edge-api-server/pkg/apis/edge"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*EdgeDevice)(nil), (*edge.EdgeDevice)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EdgeDevice_To_edge_EdgeDevice(a.(*EdgeDevice), b.(*edge.EdgeDevice), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*edge.EdgeDevice)(nil), (*EdgeDevice)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_edge_EdgeDevice_To_v1alpha1_EdgeDevice(a.(*edge.EdgeDevice), b.(*EdgeDevice), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EdgeDeviceList)(nil), (*edge.EdgeDeviceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EdgeDeviceList_To_edge_EdgeDeviceList(a.(*EdgeDeviceList), b.(*edge.EdgeDeviceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*edge.EdgeDeviceList)(nil), (*EdgeDeviceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_edge_EdgeDeviceList_To_v1alpha1_EdgeDeviceList(a.(*edge.EdgeDeviceList), b.(*EdgeDeviceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EdgeDeviceSpec)(nil), (*edge.EdgeDeviceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EdgeDeviceSpec_To_edge_EdgeDeviceSpec(a.(*EdgeDeviceSpec), b.(*edge.EdgeDeviceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*edge.EdgeDeviceSpec)(nil), (*EdgeDeviceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_edge_EdgeDeviceSpec_To_v1alpha1_EdgeDeviceSpec(a.(*edge.EdgeDeviceSpec), b.(*EdgeDeviceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EdgeNode)(nil), (*edge.EdgeNode)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EdgeNode_To_edge_EdgeNode(a.(*EdgeNode), b.(*edge.EdgeNode), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*edge.EdgeNode)(nil), (*EdgeNode)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_edge_EdgeNode_To_v1alpha1_EdgeNode(a.(*edge.EdgeNode), b.(*EdgeNode), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Workload)(nil), (*edge.Workload)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Workload_To_edge_Workload(a.(*Workload), b.(*edge.Workload), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*edge.Workload)(nil), (*Workload)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_edge_Workload_To_v1alpha1_Workload(a.(*edge.Workload), b.(*Workload), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_EdgeDevice_To_edge_EdgeDevice(in *EdgeDevice, out *edge.EdgeDevice, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_EdgeDeviceSpec_To_edge_EdgeDeviceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	out.Status = edge.EdgeDeviceStatus(in.Status)
	return nil
}

// Convert_v1alpha1_EdgeDevice_To_edge_EdgeDevice is an autogenerated conversion function.
func Convert_v1alpha1_EdgeDevice_To_edge_EdgeDevice(in *EdgeDevice, out *edge.EdgeDevice, s conversion.Scope) error {
	return autoConvert_v1alpha1_EdgeDevice_To_edge_EdgeDevice(in, out, s)
}

func autoConvert_edge_EdgeDevice_To_v1alpha1_EdgeDevice(in *edge.EdgeDevice, out *EdgeDevice, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_edge_EdgeDeviceSpec_To_v1alpha1_EdgeDeviceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	out.Status = EdgeDeviceStatus(in.Status)
	return nil
}

// Convert_edge_EdgeDevice_To_v1alpha1_EdgeDevice is an autogenerated conversion function.
func Convert_edge_EdgeDevice_To_v1alpha1_EdgeDevice(in *edge.EdgeDevice, out *EdgeDevice, s conversion.Scope) error {
	return autoConvert_edge_EdgeDevice_To_v1alpha1_EdgeDevice(in, out, s)
}

func autoConvert_v1alpha1_EdgeDeviceList_To_edge_EdgeDeviceList(in *EdgeDeviceList, out *edge.EdgeDeviceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]edge.EdgeDevice)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_EdgeDeviceList_To_edge_EdgeDeviceList is an autogenerated conversion function.
func Convert_v1alpha1_EdgeDeviceList_To_edge_EdgeDeviceList(in *EdgeDeviceList, out *edge.EdgeDeviceList, s conversion.Scope) error {
	return autoConvert_v1alpha1_EdgeDeviceList_To_edge_EdgeDeviceList(in, out, s)
}

func autoConvert_edge_EdgeDeviceList_To_v1alpha1_EdgeDeviceList(in *edge.EdgeDeviceList, out *EdgeDeviceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]EdgeDevice)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_edge_EdgeDeviceList_To_v1alpha1_EdgeDeviceList is an autogenerated conversion function.
func Convert_edge_EdgeDeviceList_To_v1alpha1_EdgeDeviceList(in *edge.EdgeDeviceList, out *EdgeDeviceList, s conversion.Scope) error {
	return autoConvert_edge_EdgeDeviceList_To_v1alpha1_EdgeDeviceList(in, out, s)
}

func autoConvert_v1alpha1_EdgeDeviceSpec_To_edge_EdgeDeviceSpec(in *EdgeDeviceSpec, out *edge.EdgeDeviceSpec, s conversion.Scope) error {
	out.Nodes = *(*edge.EdgeNodes)(unsafe.Pointer(&in.Nodes))
	return nil
}

// Convert_v1alpha1_EdgeDeviceSpec_To_edge_EdgeDeviceSpec is an autogenerated conversion function.
func Convert_v1alpha1_EdgeDeviceSpec_To_edge_EdgeDeviceSpec(in *EdgeDeviceSpec, out *edge.EdgeDeviceSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_EdgeDeviceSpec_To_edge_EdgeDeviceSpec(in, out, s)
}

func autoConvert_edge_EdgeDeviceSpec_To_v1alpha1_EdgeDeviceSpec(in *edge.EdgeDeviceSpec, out *EdgeDeviceSpec, s conversion.Scope) error {
	out.Nodes = *(*EdgeNodes)(unsafe.Pointer(&in.Nodes))
	return nil
}

// Convert_edge_EdgeDeviceSpec_To_v1alpha1_EdgeDeviceSpec is an autogenerated conversion function.
func Convert_edge_EdgeDeviceSpec_To_v1alpha1_EdgeDeviceSpec(in *edge.EdgeDeviceSpec, out *EdgeDeviceSpec, s conversion.Scope) error {
	return autoConvert_edge_EdgeDeviceSpec_To_v1alpha1_EdgeDeviceSpec(in, out, s)
}

func autoConvert_v1alpha1_EdgeNode_To_edge_EdgeNode(in *EdgeNode, out *edge.EdgeNode, s conversion.Scope) error {
	out.Name = in.Name
	out.Status = edge.EdgeNodeStatus(in.Status)
	out.LastSeenTimestamp = in.LastSeenTimestamp
	out.Workloads = *(*[]edge.Workload)(unsafe.Pointer(&in.Workloads))
	return nil
}

// Convert_v1alpha1_EdgeNode_To_edge_EdgeNode is an autogenerated conversion function.
func Convert_v1alpha1_EdgeNode_To_edge_EdgeNode(in *EdgeNode, out *edge.EdgeNode, s conversion.Scope) error {
	return autoConvert_v1alpha1_EdgeNode_To_edge_EdgeNode(in, out, s)
}

func autoConvert_edge_EdgeNode_To_v1alpha1_EdgeNode(in *edge.EdgeNode, out *EdgeNode, s conversion.Scope) error {
	out.Name = in.Name
	out.Status = EdgeNodeStatus(in.Status)
	out.LastSeenTimestamp = in.LastSeenTimestamp
	out.Workloads = *(*[]Workload)(unsafe.Pointer(&in.Workloads))
	return nil
}

// Convert_edge_EdgeNode_To_v1alpha1_EdgeNode is an autogenerated conversion function.
func Convert_edge_EdgeNode_To_v1alpha1_EdgeNode(in *edge.EdgeNode, out *EdgeNode, s conversion.Scope) error {
	return autoConvert_edge_EdgeNode_To_v1alpha1_EdgeNode(in, out, s)
}

func autoConvert_v1alpha1_Workload_To_edge_Workload(in *Workload, out *edge.Workload, s conversion.Scope) error {
	out.Name = in.Name
	out.State = in.State
	out.SubState = in.SubState
	return nil
}

// Convert_v1alpha1_Workload_To_edge_Workload is an autogenerated conversion function.
func Convert_v1alpha1_Workload_To_edge_Workload(in *Workload, out *edge.Workload, s conversion.Scope) error {
	return autoConvert_v1alpha1_Workload_To_edge_Workload(in, out, s)
}

func autoConvert_edge_Workload_To_v1alpha1_Workload(in *edge.Workload, out *Workload, s conversion.Scope) error {
	out.Name = in.Name
	out.State = in.State
	out.SubState = in.SubState
	return nil
}

// Convert_edge_Workload_To_v1alpha1_Workload is an autogenerated conversion function.
func Convert_edge_Workload_To_v1alpha1_Workload(in *edge.Workload, out *Workload, s conversion.Scope) error {
	return autoConvert_edge_Workload_To_v1alpha1_Workload(in, out, s)
}
