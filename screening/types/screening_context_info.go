// Package types implements all the types used by the Screening protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ScreeningContextInfo is a type within the Screening protocol
type ScreeningContextInfo struct {
	types.Structure
	Key   types.String
	Value types.String
}

// WriteTo writes the ScreeningContextInfo to the given writable
func (sci ScreeningContextInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sci.Key.WriteTo(contentWritable)
	sci.Value.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sci.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ScreeningContextInfo from the given readable
func (sci *ScreeningContextInfo) ExtractFrom(readable types.Readable) error {
	if err := sci.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningContextInfo header. %s", err.Error())
	}

	if err := sci.Key.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningContextInfo.Key. %s", err.Error())
	}

	if err := sci.Value.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningContextInfo.Value. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ScreeningContextInfo
func (sci ScreeningContextInfo) Copy() types.RVType {
	copied := NewScreeningContextInfo()

	copied.StructureVersion = sci.StructureVersion
	copied.Key = sci.Key
	copied.Value = sci.Value

	return copied
}

// Equals checks if the given ScreeningContextInfo contains the same data as the current ScreeningContextInfo
func (sci ScreeningContextInfo) Equals(o types.RVType) bool {
	if _, ok := o.(ScreeningContextInfo); !ok {
		return false
	}

	other := o.(ScreeningContextInfo)

	if sci.StructureVersion != other.StructureVersion {
		return false
	}

	if sci.Key != other.Key {
		return false
	}

	return sci.Value == other.Value
}

// CopyRef copies the current value of the ScreeningContextInfo
// and returns a pointer to the new copy
func (sci ScreeningContextInfo) CopyRef() types.RVTypePtr {
	copied := sci
	return &copied
}

// Deref takes a pointer to the ScreeningContextInfo
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sci *ScreeningContextInfo) Deref() types.RVType {
	return *sci
}

// String returns the string representation of the ScreeningContextInfo
func (sci ScreeningContextInfo) String() string {
	return sci.FormatToString(0)
}

// FormatToString pretty-prints the ScreeningContextInfo using the provided indentation level
func (sci ScreeningContextInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ScreeningContextInfo{\n")
	b.WriteString(fmt.Sprintf("%sKey: %s,\n", indentationValues, sci.Key))
	b.WriteString(fmt.Sprintf("%sValue: %s\n", indentationValues, sci.Value))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewScreeningContextInfo returns a new ScreeningContextInfo
func NewScreeningContextInfo() ScreeningContextInfo {
	return ScreeningContextInfo{
		Key:   types.NewString(""),
		Value: types.NewString(""),
	}
}
