package friends_wiiu_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

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

// String returns a string representation of the struct
func (friendInfo *FriendInfo) String() string {
	return friendInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (friendInfo *FriendInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, friendInfo.StructureVersion()))

	if friendInfo.NNAInfo != nil {
		b.WriteString(fmt.Sprintf("%sNNAInfo: %s,\n", indentationValues, friendInfo.NNAInfo.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sNNAInfo: nil,\n", indentationValues))
	}

	if friendInfo.Presence != nil {
		b.WriteString(fmt.Sprintf("%sPresence: %s,\n", indentationValues, friendInfo.Presence.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPresence: nil,\n", indentationValues))
	}

	if friendInfo.Status != nil {
		b.WriteString(fmt.Sprintf("%sStatus: %s,\n", indentationValues, friendInfo.Status.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sStatus: nil,\n", indentationValues))
	}

	if friendInfo.BecameFriend != nil {
		b.WriteString(fmt.Sprintf("%sBecameFriend: %s,\n", indentationValues, friendInfo.BecameFriend.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sBecameFriend: nil,\n", indentationValues))
	}

	if friendInfo.LastOnline != nil {
		b.WriteString(fmt.Sprintf("%sLastOnline: %s,\n", indentationValues, friendInfo.LastOnline.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sLastOnline: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sUnknown: %d\n", indentationValues, friendInfo.Unknown))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendInfo returns a new FriendInfo
func NewFriendInfo() *FriendInfo {
	return &FriendInfo{}
}
