package friends_3ds_types

import "github.com/PretendoNetwork/nex-go"

// FriendPersistentInfo contains user settings
type FriendPersistentInfo struct {
	nex.Structure
	PID              uint32
	Region           uint8
	Country          uint8
	Area             uint8
	Language         uint8
	Platform         uint8
	GameKey          *GameKey
	Message          string
	MessageUpdatedAt *nex.DateTime //appears to be correct, but not 100% sure.
	FriendedAt       *nex.DateTime
	LastOnline       *nex.DateTime
}

// Bytes encodes the FriendPersistentInfo and returns a byte array
func (friendPersistentInfo *FriendPersistentInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(friendPersistentInfo.PID)
	stream.WriteUInt8(friendPersistentInfo.Region)
	stream.WriteUInt8(friendPersistentInfo.Country)
	stream.WriteUInt8(friendPersistentInfo.Area)
	stream.WriteUInt8(friendPersistentInfo.Language)
	stream.WriteUInt8(friendPersistentInfo.Platform)
	stream.WriteStructure(friendPersistentInfo.GameKey)
	stream.WriteString(friendPersistentInfo.Message)
	stream.WriteDateTime(friendPersistentInfo.MessageUpdatedAt)
	stream.WriteDateTime(friendPersistentInfo.FriendedAt)
	stream.WriteDateTime(friendPersistentInfo.LastOnline)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendPersistentInfo
func (friendPersistentInfo *FriendPersistentInfo) Copy() nex.StructureInterface {
	copied := NewFriendPersistentInfo()

	copied.PID = friendPersistentInfo.PID
	copied.Region = friendPersistentInfo.Region
	copied.Country = friendPersistentInfo.Country
	copied.Area = friendPersistentInfo.Area
	copied.Language = friendPersistentInfo.Language
	copied.Platform = friendPersistentInfo.Platform
	copied.GameKey = friendPersistentInfo.GameKey.Copy().(*GameKey)
	copied.Message = friendPersistentInfo.Message
	copied.MessageUpdatedAt = friendPersistentInfo.MessageUpdatedAt.Copy()
	copied.FriendedAt = friendPersistentInfo.FriendedAt.Copy()
	copied.LastOnline = friendPersistentInfo.LastOnline.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendPersistentInfo *FriendPersistentInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendPersistentInfo)

	if friendPersistentInfo.PID != other.PID {
		return false
	}

	if friendPersistentInfo.Region != other.Region {
		return false
	}

	if friendPersistentInfo.Country != other.Country {
		return false
	}

	if friendPersistentInfo.Area != other.Area {
		return false
	}

	if friendPersistentInfo.Language != other.Language {
		return false
	}

	if friendPersistentInfo.Platform != other.Platform {
		return false
	}

	if !friendPersistentInfo.GameKey.Equals(other.GameKey) {
		return false
	}

	if friendPersistentInfo.Message != other.Message {
		return false
	}

	if !friendPersistentInfo.MessageUpdatedAt.Equals(other.MessageUpdatedAt) {
		return false
	}

	if !friendPersistentInfo.FriendedAt.Equals(other.FriendedAt) {
		return false
	}

	if !friendPersistentInfo.LastOnline.Equals(other.LastOnline) {
		return false
	}

	return true
}

// NewFriendPersistentInfo returns a new FriendPersistentInfo
func NewFriendPersistentInfo() *FriendPersistentInfo {
	return &FriendPersistentInfo{}
}
