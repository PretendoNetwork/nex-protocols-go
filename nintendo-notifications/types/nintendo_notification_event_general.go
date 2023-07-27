// Package types implements all the types used by the Nintendo Notifications protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

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

// String returns a string representation of the struct
func (nintendoNotificationEventGeneral *NintendoNotificationEventGeneral) String() string {
	return nintendoNotificationEventGeneral.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (nintendoNotificationEventGeneral *NintendoNotificationEventGeneral) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NintendoNotificationEventGeneral{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, nintendoNotificationEventGeneral.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sU32Param: %d,\n", indentationValues, nintendoNotificationEventGeneral.U32Param))
	b.WriteString(fmt.Sprintf("%sU64Param1: %d,\n", indentationValues, nintendoNotificationEventGeneral.U64Param1))
	b.WriteString(fmt.Sprintf("%sU64Param2: %d,\n", indentationValues, nintendoNotificationEventGeneral.U64Param2))
	b.WriteString(fmt.Sprintf("%sStrParam: %q\n", indentationValues, nintendoNotificationEventGeneral.StrParam))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoNotificationEventGeneral returns a new NintendoNotificationEventGeneral
func NewNintendoNotificationEventGeneral() *NintendoNotificationEventGeneral {
	return &NintendoNotificationEventGeneral{}
}
