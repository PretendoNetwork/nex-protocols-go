package friends_wiiu_types

import "github.com/PretendoNetwork/nex-go"

// FriendInfo contains information about a friend
type FriendInfo struct {
	nex.Structure
	NNAInfo      *NNAInfo
	Presence     *NintendoPresenceV2
	Status       *Comment
	BecameFriend *nex.DateTime
	LastOnline   *nex.DateTime
	Unknown      uint64
}

// Bytes encodes the FriendInfo and returns a byte array
func (friendInfo *FriendInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(friendInfo.NNAInfo)
	stream.WriteStructure(friendInfo.Presence)
	stream.WriteStructure(friendInfo.Status)
	stream.WriteDateTime(friendInfo.BecameFriend)
	stream.WriteDateTime(friendInfo.LastOnline)
	stream.WriteUInt64LE(friendInfo.Unknown)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendInfo
func (friendInfo *FriendInfo) Copy() nex.StructureInterface {
	copied := NewFriendInfo()

	copied.NNAInfo = friendInfo.NNAInfo.Copy().(*NNAInfo)
	copied.Presence = friendInfo.Presence.Copy().(*NintendoPresenceV2)
	copied.Status = friendInfo.Status.Copy().(*Comment)
	copied.BecameFriend = friendInfo.BecameFriend.Copy()
	copied.LastOnline = friendInfo.LastOnline.Copy()
	copied.Unknown = friendInfo.Unknown

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendInfo *FriendInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendInfo)

	if !friendInfo.NNAInfo.Equals(other.NNAInfo) {
		return false
	}

	if !friendInfo.Presence.Equals(other.Presence) {
		return false
	}

	if !friendInfo.Status.Equals(other.Status) {
		return false
	}

	if !friendInfo.BecameFriend.Equals(other.BecameFriend) {
		return false
	}

	if !friendInfo.LastOnline.Equals(other.LastOnline) {
		return false
	}

	if friendInfo.Unknown != other.Unknown {
		return false
	}

	return true
}

// NewFriendInfo returns a new FriendInfo
func NewFriendInfo() *FriendInfo {
	return &FriendInfo{}
}
