// Package types implements all the types used by the Matchmake Extension (Mario Kart 8) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// SimpleSearchDateTimeAttribute holds data for the Matchmake Extension (Mario Kart 8) protocol
type SimpleSearchDateTimeAttribute struct {
	types.Structure
	Unknown   *types.PrimitiveU32
	Unknown2  *types.PrimitiveU32
	Unknown3  *types.PrimitiveU32
	Unknown4  *types.PrimitiveU32
	StartTime *types.DateTime
	EndTime   *types.DateTime
}

// ExtractFrom extracts the SimpleSearchDateTimeAttribute from the given readable
func (simpleSearchDateTimeAttribute *SimpleSearchDateTimeAttribute) ExtractFrom(readable types.Readable) error {
	var err error

	if err = simpleSearchDateTimeAttribute.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read SimpleSearchDateTimeAttribute header. %s", err.Error())
	}

	err = simpleSearchDateTimeAttribute.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.Unknown from stream. %s", err.Error())
	}

	err = simpleSearchDateTimeAttribute.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.Unknown2 from stream. %s", err.Error())
	}

	err = simpleSearchDateTimeAttribute.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.Unknown3 from stream. %s", err.Error())
	}

	err = simpleSearchDateTimeAttribute.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.Unknown4 from stream. %s", err.Error())
	}

	err = simpleSearchDateTimeAttribute.StartTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.StartTime from stream. %s", err.Error())
	}

	err = simpleSearchDateTimeAttribute.EndTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.EndTime from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the SimpleSearchDateTimeAttribute to the given writable
func (simpleSearchDateTimeAttribute *SimpleSearchDateTimeAttribute) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	simpleSearchDateTimeAttribute.Unknown.WriteTo(contentWritable)
	simpleSearchDateTimeAttribute.Unknown2.WriteTo(contentWritable)
	simpleSearchDateTimeAttribute.Unknown3.WriteTo(contentWritable)
	simpleSearchDateTimeAttribute.Unknown4.WriteTo(contentWritable)
	simpleSearchDateTimeAttribute.StartTime.WriteTo(contentWritable)
	simpleSearchDateTimeAttribute.EndTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	simpleSearchDateTimeAttribute.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of SimpleSearchDateTimeAttribute
func (simpleSearchDateTimeAttribute *SimpleSearchDateTimeAttribute) Copy() types.RVType {
	copied := NewSimpleSearchDateTimeAttribute()

	copied.StructureVersion = simpleSearchDateTimeAttribute.StructureVersion

	copied.Unknown = simpleSearchDateTimeAttribute.Unknown
	copied.Unknown2 = simpleSearchDateTimeAttribute.Unknown2
	copied.Unknown3 = simpleSearchDateTimeAttribute.Unknown3
	copied.Unknown4 = simpleSearchDateTimeAttribute.Unknown4
	copied.StartTime = simpleSearchDateTimeAttribute.StartTime.Copy()
	copied.EndTime = simpleSearchDateTimeAttribute.EndTime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleSearchDateTimeAttribute *SimpleSearchDateTimeAttribute) Equals(o types.RVType) bool {
	if _, ok := o.(*SimpleSearchDateTimeAttribute); !ok {
		return false
	}

	other := o.(*SimpleSearchDateTimeAttribute)

	if simpleSearchDateTimeAttribute.StructureVersion != other.StructureVersion {
		return false
	}

	if !simpleSearchDateTimeAttribute.Unknown.Equals(other.Unknown) {
		return false
	}

	if !simpleSearchDateTimeAttribute.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !simpleSearchDateTimeAttribute.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !simpleSearchDateTimeAttribute.Unknown4.Equals(other.Unknown4) {
		return false
	}

	if !simpleSearchDateTimeAttribute.StartTime.Equals(other.StartTime) {
		return false
	}

	if !simpleSearchDateTimeAttribute.EndTime.Equals(other.EndTime) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (simpleSearchDateTimeAttribute *SimpleSearchDateTimeAttribute) String() string {
	return simpleSearchDateTimeAttribute.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (simpleSearchDateTimeAttribute *SimpleSearchDateTimeAttribute) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SimpleSearchDateTimeAttribute{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, simpleSearchDateTimeAttribute.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUnknown: %d,\n", indentationValues, simpleSearchDateTimeAttribute.Unknown))
	b.WriteString(fmt.Sprintf("%sUnknown2: %d,\n", indentationValues, simpleSearchDateTimeAttribute.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %d,\n", indentationValues, simpleSearchDateTimeAttribute.Unknown3))
	b.WriteString(fmt.Sprintf("%sUnknown4: %d,\n", indentationValues, simpleSearchDateTimeAttribute.Unknown4))

	if simpleSearchDateTimeAttribute.StartTime != nil {
		b.WriteString(fmt.Sprintf("%sStartTime: %s\n", indentationValues, simpleSearchDateTimeAttribute.StartTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sStartTime: nil\n", indentationValues))
	}

	if simpleSearchDateTimeAttribute.EndTime != nil {
		b.WriteString(fmt.Sprintf("%sEndTime: %s\n", indentationValues, simpleSearchDateTimeAttribute.EndTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sEndTime: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleSearchDateTimeAttribute returns a new SimpleSearchDateTimeAttribute
func NewSimpleSearchDateTimeAttribute() *SimpleSearchDateTimeAttribute {
	return &SimpleSearchDateTimeAttribute{}
}
