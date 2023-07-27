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
	Server *nex.Server
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	// TODO: Do something
	// This protocol doesn't seem to get requests from the client, it only sends them
	// So no handling is done for in-coming requests at the moment
}

// NewProtocol returns a new Notifications protocol
func NewProtocol(server *nex.Server) *Protocol {
	notificationsProtocol := &Protocol{Server: server}

	notificationsProtocol.Setup()

	return notificationsProtocol
}
