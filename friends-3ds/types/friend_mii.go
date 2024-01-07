// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FriendMii is a data structure used by the Friends 3DS protocol to hold information about a friends Mii
type FriendMii struct {
	types.Structure
	*types.Data
	PID        *types.PID
	Mii        *Mii
	ModifiedAt *types.DateTime
}

// WriteTo writes the Mii to the given writable
func (friendMii *FriendMii) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	friendMii.PID.WriteTo(contentWritable)
	friendMii.Mii.WriteTo(contentWritable)
	friendMii.ModifiedAt.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	friendMii.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of FriendMii
func (friendMii *FriendMii) Copy() types.RVType {
	copied := NewFriendMii()

	copied.StructureVersion = friendMii.StructureVersion

	copied.Data = friendMii.Data.Copy().(*types.Data)

	copied.PID = friendMii.PID.Copy()
	copied.Mii = friendMii.Mii.Copy().(*Mii)
	copied.ModifiedAt = friendMii.ModifiedAt.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendMii *FriendMii) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendMii); !ok {
		return false
	}

	other := o.(*FriendMii)

	if friendMii.StructureVersion != other.StructureVersion {
		return false
	}

	if !friendMii.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !friendMii.PID.Equals(other.PID) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, friendMii.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, friendMii.PID.FormatToString(indentationLevel+1)))

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
