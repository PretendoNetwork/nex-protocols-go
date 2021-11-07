package nexproto

import (
	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// NotificationsProtocolID is the protocol ID for the Nintendo Notifications protocol
	NotificationsProtocolID = 0xE

	// NotificationsMethodProcessNotificationEvent is the method ID for the method ProcessNotificationEvent
	NotificationsMethodProcessNotificationEvent = 0x1
)

// NotificationsProtocol handles the Notifications protocol
type NotificationsProtocol struct {
	server *nex.Server
}

// NotificationEvent holds general purpose notification data
type NotificationEvent struct {
	nex.Structure
	PIDSource uint32
	Type      uint32
	Param1    uint32
	Param2    uint32
	StrParam  string
}

// Bytes encodes the NotificationEvent and returns a byte array
func (notificationEventGeneral *NotificationEvent) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(notificationEventGeneral.PIDSource)
	stream.WriteUInt32LE(notificationEventGeneral.Type)
	stream.WriteUInt32LE(notificationEventGeneral.Param1)
	stream.WriteUInt32LE(notificationEventGeneral.Param2)
	stream.WriteString(notificationEventGeneral.StrParam)

	return stream.Bytes()
}

// NotificationEvent returns a new NotificationEvent
func NewNotificationEvent() *NotificationEvent {
	return &NotificationEvent{}
}

// Setup initializes the protocol
func (notificationsProtocol *NotificationsProtocol) Setup() {
	// TODO: Do something
	// This protocol doesn't seem to get requests from the client, it only sends them
	// So no handling is done for in-coming requests at the moment
}

// NewNotificationsProtocol returns a new NotificationsProtocol
func NewNotificationsProtocol(server *nex.Server) *NotificationsProtocol {
	notificationsProtocol := &NotificationsProtocol{server: server}

	notificationsProtocol.Setup()

	return notificationsProtocol
}
