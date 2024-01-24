// Package protocol implements the Nintendo Notfications protocol
package protocol

import nex "github.com/PretendoNetwork/nex-go"

const (
	// ProtocolID is the protocol ID for the Nintendo Notifications protocol
	ProtocolID = 0x64

	// MethodProcessNintendoNotificationEvent1 is the method ID for the method ProcessNintendoNotificationEvent (1)
	MethodProcessNintendoNotificationEvent1 = 0x1

	// MethodProcessNintendoNotificationEvent2 is the method ID for the method ProcessNintendoNotificationEvent (2)
	MethodProcessNintendoNotificationEvent2 = 0x2
)

// Protocol handles the NintendoNotifications protocol
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

// NewProtocol returns a new Nintendo Notifications protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	return &Protocol{server: server}
}
