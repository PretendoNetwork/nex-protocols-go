package friends_3ds_types

import "github.com/PretendoNetwork/nex-go"

type FriendMii struct {
	nex.Structure
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

	copied.PID = friendMii.PID
	copied.Mii = friendMii.Mii.Copy().(*Mii)
	copied.ModifiedAt = friendMii.ModifiedAt.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendMii *FriendMii) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendMii)

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

// NewFriendMii returns a new FriendMii
func NewFriendMii() *FriendMii {
	return &FriendMii{}
}
