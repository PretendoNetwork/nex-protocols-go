package nintendo_notifications_types

import "github.com/PretendoNetwork/nex-go"

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

// Copy returns a new copied instance of NintendoNotificationEvent
func (nintendoNotificationEvent *NintendoNotificationEvent) Copy() nex.StructureInterface {
	copied := NewNintendoNotificationEvent()

	copied.Type = nintendoNotificationEvent.Type
	copied.SenderPID = nintendoNotificationEvent.SenderPID
	copied.DataHolder = nintendoNotificationEvent.DataHolder.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nintendoNotificationEvent *NintendoNotificationEvent) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NintendoNotificationEvent)

	if nintendoNotificationEvent.Type != other.Type {
		return false
	}

	if nintendoNotificationEvent.SenderPID != other.SenderPID {
		return false
	}

	if !nintendoNotificationEvent.DataHolder.Equals(other.DataHolder) {
		return false
	}

	return true
}

// NewNintendoNotificationEvent returns a new NintendoNotificationEvent
func NewNintendoNotificationEvent() *NintendoNotificationEvent {
	return &NintendoNotificationEvent{}
}
