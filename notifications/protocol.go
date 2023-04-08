package notifications

import (
	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// ProtocolID is the protocol ID for the Nintendo Notifications protocol
	ProtocolID = 0xE

	// MethodProcessNotificationEvent is the method ID for the method ProcessNotificationEvent
	MethodProcessNotificationEvent = 0x1
)

// NotificationsProtocol handles the Notifications protocol
type NotificationsProtocol struct {
	Server *nex.Server
}

// Setup initializes the protocol
func (protocol *NotificationsProtocol) Setup() {
	// TODO: Do something
	// This protocol doesn't seem to get requests from the client, it only sends them
	// So no handling is done for in-coming requests at the moment
}

// NewNotificationsProtocol returns a new NotificationsProtocol
func NewNotificationsProtocol(server *nex.Server) *NotificationsProtocol {
	notificationsProtocol := &NotificationsProtocol{Server: server}

	notificationsProtocol.Setup()

	return notificationsProtocol
}
