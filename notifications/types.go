package notifications

import nex "github.com/PretendoNetwork/nex-go"

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
