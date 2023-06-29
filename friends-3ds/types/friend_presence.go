package friends_3ds_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// FriendPresence contains information about a users online presence
type FriendPresence struct {
	nex.Structure
	PID      uint32
	Presence *NintendoPresence
}

// Bytes encodes the FriendPresence and returns a byte array
func (presence *FriendPresence) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(presence.PID)
	stream.WriteStructure(presence.Presence)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendPresence
func (presence *FriendPresence) Copy() nex.StructureInterface {
	copied := NewFriendPresence()

	copied.PID = presence.PID
	copied.Presence = presence.Presence.Copy().(*NintendoPresence)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (presence *FriendPresence) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendPresence)

	if presence.PID != other.PID {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, presence.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, presence.PID))

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
