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
	Server nex.ServerInterface
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	// TODO: Do something
	// This protocol doesn't seem to get requests from the client, it only sends them
	// So no handling is done for in-coming requests at the moment
}

// NewProtocol returns a new Nintendo Notifications protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	nintendoNotificationsProtocol := &Protocol{Server: server}

	nintendoNotificationsProtocol.Setup()

	return nintendoNotificationsProtocol
}
