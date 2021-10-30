package nexproto

import (
	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// NintendoNotificationsProtocolID is the protocol ID for the Nintendo Notifications protocol
	NintendoNotificationsProtocolID = 0x64

	// NintendoNotificationsMethodProcessNintendoNotificationEvent1 is the method ID for the method ProcessNintendoNotificationEvent (1)
	NintendoNotificationsMethodProcessNintendoNotificationEvent1 = 0x1

	// NintendoNotificationsMethodProcessNintendoNotificationEvent2 is the method ID for the method ProcessNintendoNotificationEvent (2)
	NintendoNotificationsMethodProcessNintendoNotificationEvent2 = 0x2
)

// NintendoNotificationsProtocol handles the NintendoNotifications protocol
type NintendoNotificationsProtocol struct {
	server *nex.Server
}

// NintendoNotificationEvent is used to send data about a notification event to a client
type NintendoNotificationEvent struct {
	nex.Structure
	Type       uint32
	SenderPID  uint32
	DataHolder *nex.DataHolder
}

// Bytes encodes the NintendoNotificationEvent and returns a byte array
func (nintendoNotificationEvent *NintendoNotificationEvent) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(nintendoNotificationEvent.Type)
	stream.WriteUInt32LE(nintendoNotificationEvent.SenderPID)
	stream.WriteDataHolder(nintendoNotificationEvent.DataHolder)

	return stream.Bytes()
}

// NewNintendoNotificationEvent returns a new NintendoNotificationEvent
func NewNintendoNotificationEvent() *NintendoNotificationEvent {
	return &NintendoNotificationEvent{}
}

// Setup initializes the protocol
func (nintendoNotificationsProtocol *NintendoNotificationsProtocol) Setup() {
	// TODO: Do something
	// This protocol doesn't seem to get requests from the client, it only sends them
	// So no handling is done for in-coming requests at the moment
}

// NewNintendoNotificationsProtocol returns a new NintendoNotificationsProtocol
func NewNintendoNotificationsProtocol(server *nex.Server) *NintendoNotificationsProtocol {
	nintendoNotificationsProtocol := &NintendoNotificationsProtocol{server: server}

	nintendoNotificationsProtocol.Setup()

	return nintendoNotificationsProtocol
}
