// Package types implements all the types used by the Nintendo Notifications protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// NintendoNotificationEventGeneral holds general purpose notification data
type NintendoNotificationEventGeneral struct {
	types.Structure
	U32Param  *types.PrimitiveU32
	U64Param1 *types.PrimitiveU64
	U64Param2 *types.PrimitiveU64
	StrParam  string
}

// WriteTo writes the NintendoNotificationEventGeneral to the given writable
func (nintendoNotificationEventGeneral *NintendoNotificationEventGeneral) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nintendoNotificationEventGeneral.U32Param.WriteTo(contentWritable)
	nintendoNotificationEventGeneral.U64Param1.WriteTo(contentWritable)
	nintendoNotificationEventGeneral.U64Param2.WriteTo(contentWritable)
	nintendoNotificationEventGeneral.StrParam.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nintendoNotificationEventGeneral.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of NintendoNotificationEventGeneral
func (nintendoNotificationEventGeneral *NintendoNotificationEventGeneral) Copy() types.RVType {
	copied := NewNintendoNotificationEventGeneral()

	copied.StructureVersion = nintendoNotificationEventGeneral.StructureVersion

	copied.U32Param = nintendoNotificationEventGeneral.U32Param
	copied.U64Param1 = nintendoNotificationEventGeneral.U64Param1
	copied.U64Param2 = nintendoNotificationEventGeneral.U64Param2
	copied.StrParam = nintendoNotificationEventGeneral.StrParam

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nintendoNotificationEventGeneral *NintendoNotificationEventGeneral) Equals(o types.RVType) bool {
	if _, ok := o.(*NintendoNotificationEventGeneral); !ok {
		return false
	}

	other := o.(*NintendoNotificationEventGeneral)

	if nintendoNotificationEventGeneral.StructureVersion != other.StructureVersion {
		return false
	}

	if !nintendoNotificationEventGeneral.U32Param.Equals(other.U32Param) {
		return false
	}

	if !nintendoNotificationEventGeneral.U64Param1.Equals(other.U64Param1) {
		return false
	}

	if !nintendoNotificationEventGeneral.U64Param2.Equals(other.U64Param2) {
		return false
	}

	if !nintendoNotificationEventGeneral.StrParam.Equals(other.StrParam) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, nintendoNotificationEventGeneral.StructureVersion))
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
