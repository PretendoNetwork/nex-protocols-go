package nintendo_notifications

import nex "github.com/PretendoNetwork/nex-go"

// NintendoNotificationEventGeneral holds general purpose notification data
type NintendoNotificationEventGeneral struct {
	nex.Structure
	U32Param  uint32
	U64Param1 uint64
	U64Param2 uint64
	StrParam  string
}

// Bytes encodes the NintendoNotificationEventGeneral and returns a byte array
func (nintendoNotificationEventGeneral *NintendoNotificationEventGeneral) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(nintendoNotificationEventGeneral.U32Param)
	stream.WriteUInt64LE(nintendoNotificationEventGeneral.U64Param1)
	stream.WriteUInt64LE(nintendoNotificationEventGeneral.U64Param2)
	stream.WriteString(nintendoNotificationEventGeneral.StrParam)

	return stream.Bytes()
}

// Copy returns a new copied instance of NintendoNotificationEventGeneral
func (nintendoNotificationEventGeneral *NintendoNotificationEventGeneral) Copy() nex.StructureInterface {
	copied := NewNintendoNotificationEventGeneral()

	copied.U32Param = nintendoNotificationEventGeneral.U32Param
	copied.U64Param1 = nintendoNotificationEventGeneral.U64Param1
	copied.U64Param2 = nintendoNotificationEventGeneral.U64Param2
	copied.StrParam = nintendoNotificationEventGeneral.StrParam

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nintendoNotificationEventGeneral *NintendoNotificationEventGeneral) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NintendoNotificationEventGeneral)

	if nintendoNotificationEventGeneral.U32Param != other.U32Param {
		return false
	}

	if nintendoNotificationEventGeneral.U64Param1 != other.U64Param1 {
		return false
	}

	if nintendoNotificationEventGeneral.U64Param2 != other.U64Param2 {
		return false
	}

	if nintendoNotificationEventGeneral.StrParam != other.StrParam {
		return false
	}

	return true
}

// NintendoNotificationEventGeneral returns a new NintendoNotificationEventGeneral
func NewNintendoNotificationEventGeneral() *NintendoNotificationEventGeneral {
	return &NintendoNotificationEventGeneral{}
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
