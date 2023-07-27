// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// FriendMii is a data structure used by the Friends 3DS protocol to hold information about a friends Mii
type FriendMii struct {
	nex.Structure
	*nex.Data
	PID        uint32
	Mii        *Mii
	ModifiedAt *nex.DateTime
}

// Bytes encodes the Mii and returns a byte array
func (friendMii *FriendMii) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(friendMii.PID)
	stream.WriteStructure(friendMii.Mii)
	stream.WriteDateTime(friendMii.ModifiedAt)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendMii
func (friendMii *FriendMii) Copy() nex.StructureInterface {
	copied := NewFriendMii()

	copied.Data = friendMii.ParentType().Copy().(*nex.Data)
	copied.SetParentType(copied.Data)

	copied.PID = friendMii.PID
	copied.Mii = friendMii.Mii.Copy().(*Mii)
	copied.ModifiedAt = friendMii.ModifiedAt.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendMii *FriendMii) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendMii)

	if !friendMii.ParentType().Equals(other.ParentType()) {
		return false
	}

	if friendMii.PID != other.PID {
		return false
	}

	if !friendMii.Mii.Equals(other.Mii) {
		return false
	}

	if !friendMii.ModifiedAt.Equals(other.ModifiedAt) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (friendMii *FriendMii) String() string {
	return friendMii.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (friendMii *FriendMii) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendMii{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, friendMii.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, friendMii.PID))

	if friendMii.Mii != nil {
		b.WriteString(fmt.Sprintf("%sMii: %s,\n", indentationValues, friendMii.Mii.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sMii: nil,\n", indentationValues))
	}

	if friendMii.ModifiedAt != nil {
		b.WriteString(fmt.Sprintf("%sModifiedAt: %s\n", indentationValues, friendMii.ModifiedAt.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sModifiedAt: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendMii returns a new FriendMii
func NewFriendMii() *FriendMii {
	return &FriendMii{}
}
