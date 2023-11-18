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
	ObjectID            uint32
	OwnerPID            *nex.PID
	Attributes          []uint32
	Metadata            []byte
	CommunityIDMiiverse uint32
	CommunityCode       string
	DatetimeAttribute   *SimpleSearchDateTimeAttribute
}

// ExtractFromStream extracts a SimpleSearchObject structure from a stream
func (simpleSearchObject *SimpleSearchObject) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	simpleSearchObject.ObjectID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.ObjectID from stream. %s", err.Error())
	}

	simpleSearchObject.OwnerPID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.OwnerPID from stream. %s", err.Error())
	}

	simpleSearchObject.Attributes, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.Attributes from stream. %s", err.Error())
	}

	simpleSearchObject.Metadata, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.Metadata from stream. %s", err.Error())
	}

	simpleSearchObject.CommunityIDMiiverse, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.CommunityIDMiiverse from stream. %s", err.Error())
	}

	simpleSearchObject.CommunityCode, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.CommunityCode from stream. %s", err.Error())
	}

	simpleSearchObject.DatetimeAttribute, err = nex.StreamReadStructure(stream, NewSimpleSearchDateTimeAttribute())
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.DatetimeAttribute from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the SimpleSearchObject and returns a byte array
func (simpleSearchObject *SimpleSearchObject) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(simpleSearchObject.ObjectID)
	stream.WritePID(simpleSearchObject.OwnerPID)
	stream.WriteListUInt32LE(simpleSearchObject.Attributes)
	stream.WriteQBuffer(simpleSearchObject.Metadata)
	stream.WriteUInt32LE(simpleSearchObject.CommunityIDMiiverse)
	stream.WriteString(simpleSearchObject.CommunityCode)
	stream.WriteStructure(simpleSearchObject.DatetimeAttribute)

	return stream.Bytes()
}

// Copy returns a new copied instance of SimpleSearchObject
func (simpleSearchObject *SimpleSearchObject) Copy() nex.StructureInterface {
	copied := NewSimpleSearchObject()

	copied.SetStructureVersion(simpleSearchObject.StructureVersion())

	copied.ObjectID = simpleSearchObject.ObjectID
	copied.OwnerPID = simpleSearchObject.OwnerPID.Copy()
	copied.Attributes = make([]uint32, len(simpleSearchObject.Attributes))

	copy(copied.Attributes, simpleSearchObject.Attributes)

	copied.Metadata = simpleSearchObject.Metadata
	copied.CommunityIDMiiverse = simpleSearchObject.CommunityIDMiiverse
	copied.CommunityCode = simpleSearchObject.CommunityCode
	copied.DatetimeAttribute = simpleSearchObject.DatetimeAttribute.Copy().(*SimpleSearchDateTimeAttribute)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleSearchObject *SimpleSearchObject) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SimpleSearchObject)

	if simpleSearchObject.StructureVersion() != other.StructureVersion() {
		return false
	}

	if simpleSearchObject.ObjectID != other.ObjectID {
		return false
	}

	if !simpleSearchObject.OwnerPID.Equals(other.OwnerPID) {
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

	if simpleSearchObject.CommunityIDMiiverse != other.CommunityIDMiiverse {
		return false
	}

	if simpleSearchObject.CommunityCode != other.CommunityCode {
		return false
	}

	if !simpleSearchObject.DatetimeAttribute.Equals(other.DatetimeAttribute) {
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
	b.WriteString(fmt.Sprintf("%sObjectID: %d,\n", indentationValues, simpleSearchObject.ObjectID))
	b.WriteString(fmt.Sprintf("%sOwnerPID: %s,\n", indentationValues, simpleSearchObject.OwnerPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sAttributes: %v,\n", indentationValues, simpleSearchObject.Attributes))
	b.WriteString(fmt.Sprintf("%sMetadata: %x,\n", indentationValues, simpleSearchObject.Metadata))
	b.WriteString(fmt.Sprintf("%sCommunityIDMiiverse: %d,\n", indentationValues, simpleSearchObject.CommunityIDMiiverse))
	b.WriteString(fmt.Sprintf("%sCommunityCode: %q,\n", indentationValues, simpleSearchObject.CommunityCode))

	if simpleSearchObject.DatetimeAttribute != nil {
		b.WriteString(fmt.Sprintf("%sDatetimeAttribute: %s\n", indentationValues, simpleSearchObject.DatetimeAttribute.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sDatetimeAttribute: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleSearchObject returns a new SimpleSearchObject
func NewSimpleSearchObject() *SimpleSearchObject {
	return &SimpleSearchObject{}
}
