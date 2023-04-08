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

// NewNintendoNotificationEvent returns a new NintendoNotificationEvent
func NewNintendoNotificationEvent() *NintendoNotificationEvent {
	return &NintendoNotificationEvent{}
}
