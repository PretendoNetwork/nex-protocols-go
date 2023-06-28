package notifications_types

import "github.com/PretendoNetwork/nex-go"

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
func (notificationEvent *NotificationEvent) Bytes(stream *nex.StreamOut) []byte {
	nexVersion := stream.Server.NEXVersion()

	stream.WriteUInt32LE(notificationEvent.PIDSource)
	stream.WriteUInt32LE(notificationEvent.Type)
	stream.WriteUInt32LE(notificationEvent.Param1)
	stream.WriteUInt32LE(notificationEvent.Param2)
	stream.WriteString(notificationEvent.StrParam)

	if nexVersion.Major >= 3 && nexVersion.Minor >= 4 {
		stream.WriteUInt32LE(notificationEvent.Param3)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of NotificationEvent
func (notificationEvent *NotificationEvent) Copy() nex.StructureInterface {
	copied := NewNotificationEvent()

	copied.PIDSource = notificationEvent.PIDSource
	copied.Type = notificationEvent.Type
	copied.Param1 = notificationEvent.Param1
	copied.Param2 = notificationEvent.Param2
	copied.StrParam = notificationEvent.StrParam
	copied.Param3 = notificationEvent.Param3

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (notificationEvent *NotificationEvent) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NotificationEvent)

	if notificationEvent.PIDSource != other.PIDSource {
		return false
	}

	if notificationEvent.Type != other.Type {
		return false
	}

	if notificationEvent.Param1 != other.Param1 {
		return false
	}

	if notificationEvent.Param2 != other.Param2 {
		return false
	}

	if notificationEvent.StrParam != other.StrParam {
		return false
	}

	if notificationEvent.Param3 != other.Param3 {
		return false
	}

	return true
}

// NotificationEvent returns a new NotificationEvent
func NewNotificationEvent() *NotificationEvent {
	return &NotificationEvent{}
}
