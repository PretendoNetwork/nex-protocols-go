// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FriendRelationship contains information about a users relationship with another PID
type FriendRelationship struct {
	types.Structure
	*types.Data
	PID              *types.PID
	LFC              *types.PrimitiveU64
	RelationshipType *types.PrimitiveU8
}

// WriteTo writes the FriendRelationship to the given writable
func (relationship *FriendRelationship) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	relationship.PID.WriteTo(contentWritable)
	relationship.LFC.WriteTo(contentWritable)
	relationship.RelationshipType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	relationship.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of FriendRelationship
func (relationship *FriendRelationship) Copy() types.RVType {
	copied := NewFriendRelationship()

	copied.StructureVersion = relationship.StructureVersion

	copied.Data = relationship.Data.Copy().(*types.Data)

	copied.PID = relationship.PID.Copy()
	copied.LFC = relationship.LFC
	copied.RelationshipType = relationship.RelationshipType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (relationship *FriendRelationship) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendRelationship); !ok {
		return false
	}

	other := o.(*FriendRelationship)

	if relationship.StructureVersion != other.StructureVersion {
		return false
	}

	if !relationship.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !relationship.PID.Equals(other.PID) {
		return false
	}

	if !relationship.LFC.Equals(other.LFC) {
		return false
	}

	if !relationship.RelationshipType.Equals(other.RelationshipType) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (relationship *FriendRelationship) String() string {
	return relationship.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (relationship *FriendRelationship) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendRelationship{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, relationship.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, relationship.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sLFC: %d,\n", indentationValues, relationship.LFC))
	b.WriteString(fmt.Sprintf("%sRelationshipType: %d\n", indentationValues, relationship.RelationshipType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendRelationship returns a new FriendRelationship
func NewFriendRelationship() *FriendRelationship {
	return &FriendRelationship{}
}
