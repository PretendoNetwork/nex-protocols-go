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
	Unknown  uint32
	Unknown2 uint32
	Unknown3 uint32
	Unknown4 uint32
	Start    *nex.DateTime
	End      *nex.DateTime
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

	simpleSearchDateTimeAttribute.Start, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.Start from stream. %s", err.Error())
	}

	simpleSearchDateTimeAttribute.End, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchDateTimeAttribute.End from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the SimpleSearchDateTimeAttribute and returns a byte array
func (simpleSearchDateTimeAttribute *SimpleSearchDateTimeAttribute) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(simpleSearchDateTimeAttribute.Unknown)
	stream.WriteUInt32LE(simpleSearchDateTimeAttribute.Unknown2)
	stream.WriteUInt32LE(simpleSearchDateTimeAttribute.Unknown3)
	stream.WriteUInt32LE(simpleSearchDateTimeAttribute.Unknown4)
	stream.WriteDateTime(simpleSearchDateTimeAttribute.Start)
	stream.WriteDateTime(simpleSearchDateTimeAttribute.End)

	return stream.Bytes()
}

// Copy returns a new copied instance of SimpleSearchDateTimeAttribute
func (simpleSearchDateTimeAttribute *SimpleSearchDateTimeAttribute) Copy() nex.StructureInterface {
	copied := NewSimpleSearchDateTimeAttribute()

	copied.Unknown = simpleSearchDateTimeAttribute.Unknown
	copied.Unknown2 = simpleSearchDateTimeAttribute.Unknown2
	copied.Unknown3 = simpleSearchDateTimeAttribute.Unknown3
	copied.Unknown4 = simpleSearchDateTimeAttribute.Unknown4
	copied.Start = simpleSearchDateTimeAttribute.Start.Copy()
	copied.End = simpleSearchDateTimeAttribute.End.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleSearchDateTimeAttribute *SimpleSearchDateTimeAttribute) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SimpleSearchDateTimeAttribute)

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

	if !simpleSearchDateTimeAttribute.Start.Equals(other.Start) {
		return false
	}

	if !simpleSearchDateTimeAttribute.End.Equals(other.End) {
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

	if simpleSearchDateTimeAttribute.Start != nil {
		b.WriteString(fmt.Sprintf("%sStart: %s\n", indentationValues, simpleSearchDateTimeAttribute.Start.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sStart: nil\n", indentationValues))
	}

	if simpleSearchDateTimeAttribute.End != nil {
		b.WriteString(fmt.Sprintf("%sEnd: %s\n", indentationValues, simpleSearchDateTimeAttribute.End.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sEnd: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleSearchDateTimeAttribute returns a new SimpleSearchDateTimeAttribute
func NewSimpleSearchDateTimeAttribute() *SimpleSearchDateTimeAttribute {
	return &SimpleSearchDateTimeAttribute{}
}
