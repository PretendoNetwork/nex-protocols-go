// Package types implements all the types used by the Matchmake Extension (Mario Kart 8) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// SimpleSearchObject holds data for the Matchmake Extension (Mario Kart 8) protocol
type SimpleSearchObject struct {
	nex.Structure
	Object     uint32
	Owner      uint32
	Attributes []uint32
	Metadata   []byte
	Community  uint32
	Community2 string
	Datetime   *SimpleSearchDateTimeAttribute
}

// ExtractFromStream extracts a SimpleSearchObject structure from a stream
func (simpleSearchObject *SimpleSearchObject) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	simpleSearchObject.Object, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.Object from stream. %s", err.Error())
	}

	simpleSearchObject.Owner, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.Owner from stream. %s", err.Error())
	}

	simpleSearchObject.Attributes, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.Attributes from stream. %s", err.Error())
	}

	simpleSearchObject.Metadata, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.Metadata from stream. %s", err.Error())
	}

	simpleSearchObject.Community, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.Community from stream. %s", err.Error())
	}

	simpleSearchObject.Community2, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.Community2 from stream. %s", err.Error())
	}

	datetime, err := stream.ReadStructure(NewSimpleSearchDateTimeAttribute())
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.Datetime from stream. %s", err.Error())
	}

	simpleSearchObject.Datetime = datetime.(*SimpleSearchDateTimeAttribute)

	return nil
}

// Bytes encodes the SimpleSearchObject and returns a byte array
func (simpleSearchObject *SimpleSearchObject) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(simpleSearchObject.Object)
	stream.WriteUInt32LE(simpleSearchObject.Owner)
	stream.WriteListUInt32LE(simpleSearchObject.Attributes)
	stream.WriteQBuffer(simpleSearchObject.Metadata)
	stream.WriteUInt32LE(simpleSearchObject.Community)
	stream.WriteString(simpleSearchObject.Community2)
	stream.WriteStructure(simpleSearchObject.Datetime)

	return stream.Bytes()
}

// Copy returns a new copied instance of SimpleSearchObject
func (simpleSearchObject *SimpleSearchObject) Copy() nex.StructureInterface {
	copied := NewSimpleSearchObject()

	copied.Object = simpleSearchObject.Object
	copied.Owner = simpleSearchObject.Owner
	copied.Attributes = make([]uint32, len(simpleSearchObject.Attributes))

	copy(copied.Attributes, simpleSearchObject.Attributes)

	copied.Metadata = simpleSearchObject.Metadata
	copied.Community = simpleSearchObject.Community
	copied.Community2 = simpleSearchObject.Community2
	copied.Datetime = simpleSearchObject.Datetime.Copy().(*SimpleSearchDateTimeAttribute)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleSearchObject *SimpleSearchObject) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SimpleSearchObject)

	if simpleSearchObject.Object != other.Object {
		return false
	}

	if simpleSearchObject.Owner != other.Owner {
		return false
	}

	if len(simpleSearchObject.Attributes) != len(other.Attributes) {
		return false
	}

	for i := 0; i < len(simpleSearchObject.Attributes); i++ {
		if simpleSearchObject.Attributes[i] != other.Attributes[i] {
			return false
		}
	}

	if !bytes.Equal(simpleSearchObject.Metadata, other.Metadata) {
		return false
	}

	if simpleSearchObject.Community != other.Community {
		return false
	}

	if simpleSearchObject.Community2 != other.Community2 {
		return false
	}

	if !simpleSearchObject.Datetime.Equals(other.Datetime) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (simpleSearchObject *SimpleSearchObject) String() string {
	return simpleSearchObject.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (simpleSearchObject *SimpleSearchObject) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SimpleSearchObject{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, simpleSearchObject.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sObject: %d,\n", indentationValues, simpleSearchObject.Object))
	b.WriteString(fmt.Sprintf("%sOwner: %d,\n", indentationValues, simpleSearchObject.Owner))
	b.WriteString(fmt.Sprintf("%sAttributes: %v,\n", indentationValues, simpleSearchObject.Attributes))
	b.WriteString(fmt.Sprintf("%sMetadata: %x,\n", indentationValues, simpleSearchObject.Metadata))
	b.WriteString(fmt.Sprintf("%sCommunity: %d,\n", indentationValues, simpleSearchObject.Community))
	b.WriteString(fmt.Sprintf("%sCommunity2: %q,\n", indentationValues, simpleSearchObject.Community2))

	if simpleSearchObject.Datetime != nil {
		b.WriteString(fmt.Sprintf("%sDatetime: %s\n", indentationValues, simpleSearchObject.Datetime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sDatetime: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleSearchObject returns a new SimpleSearchObject
func NewSimpleSearchObject() *SimpleSearchObject {
	return &SimpleSearchObject{}
}
