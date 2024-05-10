/* SPDX-License-Identifier: LGPL-2.1-or-later */

package bridges

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type MQTTClient struct {
	mqtt.Client
	opts *mqtt.ClientOptions

	ClientID uuid.UUID
}

func NewMQTTClient(broker string, port uint16, user string, password string) *MQTTClient {
	clientID := uuid.New()

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(clientID.String())
	opts.SetUsername(user)
	opts.SetPassword(password)

	return &MQTTClient{
		Client: nil,
		opts:   opts,

		ClientID: clientID,
	}
}

func (c *MQTTClient) WithConnectHandler(handler mqtt.OnConnectHandler) {
	if c.Client == nil {
		c.opts.OnConnect = handler
	}
}

func (c *MQTTClient) WithConnectionLostHandler(handler mqtt.ConnectionLostHandler) {
	if c.Client == nil {
		c.opts.OnConnectionLost = handler
	}
}

func (c *MQTTClient) WithDefaultHandler(handler mqtt.MessageHandler) {
	if c.Client == nil {
		c.opts.SetDefaultPublishHandler(handler)
	}
}

func (c *MQTTClient) Connect() error {
	c.Client = mqtt.NewClient(c.opts)

	if token := c.Client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *MQTTClient) IsConnected() bool {
	return c.Client != nil && c.Client.IsConnected()
}
