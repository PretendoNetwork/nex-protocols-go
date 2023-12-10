// Package protocol implements the Notfications protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// ProtocolID is the protocol ID for the Nintendo Notifications protocol
	ProtocolID = 0xE

	// MethodProcessNotificationEvent is the method ID for the method ProcessNotificationEvent
	MethodProcessNotificationEvent = 0x1
)

// Protocol handles the Notifications protocol
type Protocol struct {
	server nex.ServerInterface
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	// TODO: Do something
	// This protocol doesn't seem to get requests from the client, it only sends them
	// So no handling is done for in-coming requests at the moment
}

// NewProtocol returns a new Notifications protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	notificationsProtocol := &Protocol{server: server}

	notificationsProtocol.Setup()

	return notificationsProtocol
}
