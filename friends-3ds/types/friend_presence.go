// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FriendPresence contains information about a users online presence
type FriendPresence struct {
	types.Structure
	*types.Data
	PID      *types.PID
	Presence *NintendoPresence
}

// WriteTo writes the FriendPresence to the given writable
func (presence *FriendPresence) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	presence.PID.WriteTo(contentWritable)
	presence.Presence.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	presence.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of FriendPresence
func (presence *FriendPresence) Copy() types.RVType {
	copied := NewFriendPresence()

	copied.StructureVersion = presence.StructureVersion

	copied.Data = presence.Data.Copy().(*types.Data)

	copied.PID = presence.PID.Copy()
	copied.Presence = presence.Presence.Copy().(*NintendoPresence)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (presence *FriendPresence) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendPresence); !ok {
		return false
	}

	other := o.(*FriendPresence)

	if presence.StructureVersion != other.StructureVersion {
		return false
	}

	if !presence.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !presence.PID.Equals(other.PID) {
		return false
	}

	if !presence.Presence.Equals(other.Presence) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (presence *FriendPresence) String() string {
	return presence.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (presence *FriendPresence) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendPresence{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, presence.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, presence.PID.FormatToString(indentationLevel+1)))

	if presence.Presence != nil {
		b.WriteString(fmt.Sprintf("%sPresence: %s\n", indentationValues, presence.Presence.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPresence: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendPresence returns a new FriendPresence
func NewFriendPresence() *FriendPresence {
	return &FriendPresence{}
}
