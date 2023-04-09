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
	Param3    uint32
}

// Bytes encodes the NotificationEvent and returns a byte array
func (notificationEventGeneral *NotificationEvent) Bytes(stream *nex.StreamOut) []byte {
	nexVersion := stream.Server.NEXVersion()

	stream.WriteUInt32LE(notificationEventGeneral.PIDSource)
	stream.WriteUInt32LE(notificationEventGeneral.Type)
	stream.WriteUInt32LE(notificationEventGeneral.Param1)
	stream.WriteUInt32LE(notificationEventGeneral.Param2)
	stream.WriteString(notificationEventGeneral.StrParam)

	if nexVersion.Major >= 3 && nexVersion.Minor >= 5 {
		stream.WriteUInt32LE(notificationEventGeneral.Param3)
	}

	return stream.Bytes()
}

// NotificationEvent returns a new NotificationEvent
func NewNotificationEvent() *NotificationEvent {
	return &NotificationEvent{}
}
