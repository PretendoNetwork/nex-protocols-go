package nintendo_notifications_types

import "github.com/PretendoNetwork/nex-go"

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
