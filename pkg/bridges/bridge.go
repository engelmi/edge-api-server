/* SPDX-License-Identifier: LGPL-2.1-or-later */

package bridges

import (
	"context"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/engelmi/edge-api-server/pkg/apis/edge/v1alpha1"
	edge_v1alpha1 "github.com/engelmi/edge-api-server/pkg/generated/clientset/versioned/typed/edge/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

const TopicBase = "redhat/edge/device/"
const TopicRegister = TopicBase + "register"

type MQTTBridge struct {
	ctx       context.Context
	namespace string

	mqttClient *MQTTClient
	edgeClient *edge_v1alpha1.EdgeV1alpha1Client
}

func NewMQTTBridge(
	ctx context.Context,
	namespace string,
	mqttClient *MQTTClient,
	edgeClient *edge_v1alpha1.EdgeV1alpha1Client,
) *MQTTBridge {
	bridge := &MQTTBridge{
		ctx:       ctx,
		namespace: namespace,

		mqttClient: mqttClient,
		edgeClient: edgeClient,
	}

	mqttClient.WithConnectHandler(bridge.handleMQTTConnect)
	mqttClient.WithConnectionLostHandler(func(c mqtt.Client, err error) {
		klog.Errorf("Lost Connection!\n%v", err)
	})
	mqttClient.WithDefaultHandler(func(c mqtt.Client, m mqtt.Message) {
		klog.Warningf("Received message for unknown subscription: %s", m.Topic())
	})

	return bridge
}

func (bridge *MQTTBridge) Run(stopCh <-chan struct{}) error {

	if err := bridge.mqttClient.Connect(); err != nil {
		return fmt.Errorf("failed to connect to MQTT broker: %v", err)
	}

	<-stopCh
	return nil
}

func (bridge *MQTTBridge) handleMQTTConnect(c mqtt.Client) {
	fmt.Printf("Connected!\n")

	if token := bridge.mqttClient.Subscribe(TopicRegister, 0, bridge.handleEdgeDeviceRegister); token.Wait() && token.Error() != nil {
		klog.Errorf("Failed to subscribe to %s: %v", TopicRegister, token.Error())
		return
	}
}

func (bridge *MQTTBridge) handleEdgeDeviceRegister(client mqtt.Client, msg mqtt.Message) {
	req, err := Unmarshal[RegisterRequest](msg.Payload())
	if err != nil {
		klog.Errorf("Failed to unmarshal register message: %v", err)
		return
	}
	klog.Infof("Received register message for device with ID '%s'", req.DeviceID)

	topicDeviceBase := fmt.Sprintf("%s/%s", TopicBase, req.DeviceID)
	topicDeviceRegister := fmt.Sprintf("%s/register", topicDeviceBase)

	registerResult := "success"
	_, err = bridge.edgeClient.EdgeDevices(bridge.namespace).Create(bridge.ctx, &v1alpha1.EdgeDevice{
		TypeMeta: v1.TypeMeta{
			Kind:       "EdgeDevice",
			APIVersion: "org.redhat.edgeapi/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:              req.DeviceID,
			Namespace:         bridge.namespace,
			CreationTimestamp: v1.Now(),
		},
		Spec: v1alpha1.EdgeDeviceSpec{
			Nodes: v1alpha1.EdgeNodes{},
		},
		Status: v1alpha1.EdgeDeviceStatusDown,
	}, v1.CreateOptions{})
	if err != nil {
		klog.Errorf("Failed to create EdgeDevice for ID '%s': %s", req.DeviceID, err)
		registerResult = "failed"
	}

	resp, err := Marshal(RegisterResponse{
		Result: registerResult,
	})
	if err != nil {
		klog.Errorf("Failed to create register response for ID '%s': %v", req.DeviceID, err)
		return
	}

	topicDeviceNodes := fmt.Sprintf("%s/update", topicDeviceBase)
	if token := bridge.mqttClient.Subscribe(topicDeviceNodes, 0, bridge.handleEdgeDeviceUpdate); token.Wait() && token.Error() != nil {
		klog.Errorf("Failed to subscribe to %s: %v", topicDeviceNodes, token.Error())
		return
	}

	if token := bridge.mqttClient.Publish(topicDeviceRegister, 0, false, resp); token.Wait() && token.Error() != nil {
		klog.Errorf("Failed to send register result: %s", err)
	}
}

func (bridge *MQTTBridge) handleEdgeDeviceUpdate(client mqtt.Client, msg mqtt.Message) {
	req, err := Unmarshal[DeviceUpdateRequest](msg.Payload())
	if err != nil {
		klog.Errorf("Failed to unmarshal nodes update message: %v", err)
		return
	}
	klog.Infof("Received nodes update message for device with ID '%s'", req.ID)

	device, err := bridge.edgeClient.EdgeDevices(bridge.namespace).Get(bridge.ctx, req.ID, v1.GetOptions{})
	if err != nil {
		klog.Errorf("Failed to fetch edge device with ID '%s': %v", req.ID, err)
		return
	}

	for _, reqNode := range req.Nodes {
		var node *v1alpha1.EdgeNode = nil
		for i, specNode := range device.Spec.Nodes {
			if specNode.Name == reqNode.Name {
				device.Spec.Nodes[i].Status = v1alpha1.EdgeNodeStatus(reqNode.Status)
				device.Spec.Nodes[i].LastSeenTimestamp = reqNode.LastSeenTimestamp
				node = &device.Spec.Nodes[i]
				break
			}
		}
		if node == nil {
			device.Spec.Nodes = append(device.Spec.Nodes, v1alpha1.EdgeNode{
				Name:              reqNode.Name,
				Status:            v1alpha1.EdgeNodeStatus(reqNode.Status),
				LastSeenTimestamp: reqNode.LastSeenTimestamp,
			})
			node = &device.Spec.Nodes[len(device.Spec.Nodes)-1]
		}

		for _, reqWorkload := range reqNode.Workloads {
			hasWorkload := false
			for j, workload := range node.Workloads {
				if workload.Name == reqWorkload.Name {
					node.Workloads[j].State = reqWorkload.State
					node.Workloads[j].SubState = reqWorkload.SubState
					hasWorkload = true
				}
			}
			if !hasWorkload {
				newWorkload := v1alpha1.Workload{
					Name:     reqWorkload.Name,
					State:    reqWorkload.State,
					SubState: reqWorkload.SubState,
				}
				node.Workloads = append(node.Workloads, newWorkload)
			}
		}
	}

	// update edge system status
	numberOfOnlineNodes := 0
	for _, node := range device.Spec.Nodes {
		if node.Status == "online" {
			numberOfOnlineNodes += 1
		}
	}
	if numberOfOnlineNodes == 0 {
		device.Status = "down"
	} else if numberOfOnlineNodes == len(device.Spec.Nodes) {
		device.Status = "up"
	} else {
		device.Status = "degraded"
	}

	_, err = bridge.edgeClient.EdgeDevices(bridge.namespace).Update(bridge.ctx, device, v1.UpdateOptions{})
	if err != nil {
		klog.Errorf("Failed to update nodes of edge device with ID '%s': %v", req.ID, err)
		return
	}
}
