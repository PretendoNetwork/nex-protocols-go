// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// FriendRelationship contains information about a users relationship with another PID
type FriendRelationship struct {
	nex.Structure
	*nex.Data
	PID              uint32
	LFC              uint64
	RelationshipType uint8
}

// Bytes encodes the FriendRelationship and returns a byte array
func (relationship *FriendRelationship) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(relationship.PID)
	stream.WriteUInt64LE(relationship.LFC)
	stream.WriteUInt8(relationship.RelationshipType)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendRelationship
func (relationship *FriendRelationship) Copy() nex.StructureInterface {
	copied := NewFriendRelationship()

	copied.SetStructureVersion(relationship.StructureVersion())

	if relationship.ParentType() != nil {
		copied.Data = relationship.ParentType().Copy().(*nex.Data)
	} else {
		copied.Data = nex.NewData()
	}

	copied.PID = relationship.PID
	copied.LFC = relationship.LFC
	copied.RelationshipType = relationship.RelationshipType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (relationship *FriendRelationship) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendRelationship)

	if relationship.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !relationship.ParentType().Equals(other.ParentType()) {
		return false
	}

	if relationship.PID != other.PID {
		return false
	}

	if relationship.LFC != other.LFC {
		return false
	}

	if relationship.RelationshipType != other.RelationshipType {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, relationship.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, relationship.PID))
	b.WriteString(fmt.Sprintf("%sLFC: %d,\n", indentationValues, relationship.LFC))
	b.WriteString(fmt.Sprintf("%sRelationshipType: %d\n", indentationValues, relationship.RelationshipType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendRelationship returns a new FriendRelationship
func NewFriendRelationship() *FriendRelationship {
	return &FriendRelationship{}
}
