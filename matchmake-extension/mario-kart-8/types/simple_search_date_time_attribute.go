// Package types implements all the types used by the Matchmake Extension (Mario Kart 8) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// SimpleSearchDateTimeAttribute holds data for the Matchmake Extension (Mario Kart 8) protocol
type SimpleSearchDateTimeAttribute struct {
	nex.Structure
	Unknown   uint32
	Unknown2  uint32
	Unknown3  uint32
	Unknown4  uint32
	StartTime *nex.DateTime
	EndTime   *nex.DateTime
}

// ExtractFromStream extracts a SimpleSearchDateTimeAttribute structure from a stream
func (simpleSearchDateTimeAttribute *SimpleSearchDateTimeAttribute) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	simpleSearchDateTimeAttribute.Unknown, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.Unknown from stream. %s", err.Error())
	}

	simpleSearchDateTimeAttribute.Unknown2, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.Unknown2 from stream. %s", err.Error())
	}

	simpleSearchDateTimeAttribute.Unknown3, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.Unknown3 from stream. %s", err.Error())
	}

	simpleSearchDateTimeAttribute.Unknown4, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.Unknown4 from stream. %s", err.Error())
	}

	simpleSearchDateTimeAttribute.StartTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.StartTime from stream. %s", err.Error())
	}

	simpleSearchDateTimeAttribute.EndTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.EndTime from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the SimpleSearchDateTimeAttribute and returns a byte array
func (simpleSearchDateTimeAttribute *SimpleSearchDateTimeAttribute) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(simpleSearchDateTimeAttribute.Unknown)
	stream.WriteUInt32LE(simpleSearchDateTimeAttribute.Unknown2)
	stream.WriteUInt32LE(simpleSearchDateTimeAttribute.Unknown3)
	stream.WriteUInt32LE(simpleSearchDateTimeAttribute.Unknown4)
	stream.WriteDateTime(simpleSearchDateTimeAttribute.StartTime)
	stream.WriteDateTime(simpleSearchDateTimeAttribute.EndTime)

	return stream.Bytes()
}

// Copy returns a new copied instance of SimpleSearchDateTimeAttribute
func (simpleSearchDateTimeAttribute *SimpleSearchDateTimeAttribute) Copy() nex.StructureInterface {
	copied := NewSimpleSearchDateTimeAttribute()

	copied.SetStructureVersion(simpleSearchDateTimeAttribute.StructureVersion())

	copied.Unknown = simpleSearchDateTimeAttribute.Unknown
	copied.Unknown2 = simpleSearchDateTimeAttribute.Unknown2
	copied.Unknown3 = simpleSearchDateTimeAttribute.Unknown3
	copied.Unknown4 = simpleSearchDateTimeAttribute.Unknown4
	copied.StartTime = simpleSearchDateTimeAttribute.StartTime.Copy()
	copied.EndTime = simpleSearchDateTimeAttribute.EndTime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleSearchDateTimeAttribute *SimpleSearchDateTimeAttribute) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SimpleSearchDateTimeAttribute)

	if simpleSearchDateTimeAttribute.StructureVersion() != other.StructureVersion() {
		return false
	}

	if simpleSearchDateTimeAttribute.Unknown != other.Unknown {
		return false
	}

	if simpleSearchDateTimeAttribute.Unknown2 != other.Unknown2 {
		return false
	}

	if simpleSearchDateTimeAttribute.Unknown3 != other.Unknown3 {
		return false
	}

	if simpleSearchDateTimeAttribute.Unknown4 != other.Unknown4 {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, simpleSearchDateTimeAttribute.StructureVersion()))
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
