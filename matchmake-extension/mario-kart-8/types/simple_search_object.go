// Package types implements all the types used by the Matchmake Extension (Mario Kart 8) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// SimpleSearchObject holds data for the Matchmake Extension (Mario Kart 8) protocol
type SimpleSearchObject struct {
	types.Structure
	ObjectID            *types.PrimitiveU32
	OwnerPID            *types.PID
	Attributes          *types.List[*types.PrimitiveU32]
	Metadata            []byte
	CommunityIDMiiverse *types.PrimitiveU32
	CommunityCode       string
	DatetimeAttribute   *SimpleSearchDateTimeAttribute
}

// ExtractFrom extracts the SimpleSearchObject from the given readable
func (simpleSearchObject *SimpleSearchObject) ExtractFrom(readable types.Readable) error {
	var err error

	if err = simpleSearchObject.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read SimpleSearchObject header. %s", err.Error())
	}

	err = simpleSearchObject.ObjectID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.ObjectID from stream. %s", err.Error())
	}

	err = simpleSearchObject.OwnerPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.OwnerPID from stream. %s", err.Error())
	}

	err = simpleSearchObject.Attributes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.Attributes from stream. %s", err.Error())
	}

	simpleSearchObject.Metadata, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.Metadata from stream. %s", err.Error())
	}

	err = simpleSearchObject.CommunityIDMiiverse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.CommunityIDMiiverse from stream. %s", err.Error())
	}

	err = simpleSearchObject.CommunityCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.CommunityCode from stream. %s", err.Error())
	}

	err = simpleSearchObject.DatetimeAttribute.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchObject.DatetimeAttribute from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the SimpleSearchObject to the given writable
func (simpleSearchObject *SimpleSearchObject) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	simpleSearchObject.ObjectID.WriteTo(contentWritable)
	simpleSearchObject.OwnerPID.WriteTo(contentWritable)
	simpleSearchObject.Attributes.WriteTo(contentWritable)
	stream.WriteQBuffer(simpleSearchObject.Metadata)
	simpleSearchObject.CommunityIDMiiverse.WriteTo(contentWritable)
	simpleSearchObject.CommunityCode.WriteTo(contentWritable)
	simpleSearchObject.DatetimeAttribute.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	simpleSearchObject.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of SimpleSearchObject
func (simpleSearchObject *SimpleSearchObject) Copy() types.RVType {
	copied := NewSimpleSearchObject()

	copied.StructureVersion = simpleSearchObject.StructureVersion

	copied.ObjectID = simpleSearchObject.ObjectID
	copied.OwnerPID = simpleSearchObject.OwnerPID.Copy()
	copied.Attributes = make(*types.List[*types.PrimitiveU32], len(simpleSearchObject.Attributes))

	copy(copied.Attributes, simpleSearchObject.Attributes)

	copied.Metadata = simpleSearchObject.Metadata
	copied.CommunityIDMiiverse = simpleSearchObject.CommunityIDMiiverse
	copied.CommunityCode = simpleSearchObject.CommunityCode
	copied.DatetimeAttribute = simpleSearchObject.DatetimeAttribute.Copy().(*SimpleSearchDateTimeAttribute)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleSearchObject *SimpleSearchObject) Equals(o types.RVType) bool {
	if _, ok := o.(*SimpleSearchObject); !ok {
		return false
	}

	other := o.(*SimpleSearchObject)

	if simpleSearchObject.StructureVersion != other.StructureVersion {
		return false
	}

	if !simpleSearchObject.ObjectID.Equals(other.ObjectID) {
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

	if !simpleSearchObject.Metadata.Equals(other.Metadata) {
		return false
	}

	if !simpleSearchObject.CommunityIDMiiverse.Equals(other.CommunityIDMiiverse) {
		return false
	}

	if !simpleSearchObject.CommunityCode.Equals(other.CommunityCode) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, simpleSearchObject.StructureVersion))
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
