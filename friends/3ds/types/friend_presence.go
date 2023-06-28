package friends_3ds_types

import "github.com/PretendoNetwork/nex-go"

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

// NewFriendPresence returns a new FriendPresence
func NewFriendPresence() *FriendPresence {
	return &FriendPresence{}
}
