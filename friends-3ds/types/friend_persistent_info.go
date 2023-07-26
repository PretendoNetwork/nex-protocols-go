// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// FriendPersistentInfo contains user settings
type FriendPersistentInfo struct {
	nex.Structure
	*nex.Data
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

	copied.Data = friendPersistentInfo.ParentType().Copy().(*nex.Data)
	copied.SetParentType(copied.Data)

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

	if !friendPersistentInfo.ParentType().Equals(other.ParentType()) {
		return false
	}

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

// String returns a string representation of the struct
func (friendPersistentInfo *FriendPersistentInfo) String() string {
	return friendPersistentInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (friendPersistentInfo *FriendPersistentInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendPersistentInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, friendPersistentInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, friendPersistentInfo.PID))
	b.WriteString(fmt.Sprintf("%sRegion: %d,\n", indentationValues, friendPersistentInfo.Region))
	b.WriteString(fmt.Sprintf("%sCountry: %d,\n", indentationValues, friendPersistentInfo.Country))
	b.WriteString(fmt.Sprintf("%sArea: %d,\n", indentationValues, friendPersistentInfo.Area))
	b.WriteString(fmt.Sprintf("%sLanguage: %d,\n", indentationValues, friendPersistentInfo.Language))
	b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, friendPersistentInfo.Platform))

	if friendPersistentInfo.GameKey != nil {
		b.WriteString(fmt.Sprintf("%sGameKey: %s,\n", indentationValues, friendPersistentInfo.GameKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sGameKey: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sMessage: %q,\n", indentationValues, friendPersistentInfo.Message))

	if friendPersistentInfo.MessageUpdatedAt != nil {
		b.WriteString(fmt.Sprintf("%sMessageUpdatedAt: %s,\n", indentationValues, friendPersistentInfo.MessageUpdatedAt.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sMessageUpdatedAt: nil,\n", indentationValues))
	}

	if friendPersistentInfo.FriendedAt != nil {
		b.WriteString(fmt.Sprintf("%sFriendedAt: %s,\n", indentationValues, friendPersistentInfo.FriendedAt.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sFriendedAt: nil,\n", indentationValues))
	}

	if friendPersistentInfo.LastOnline != nil {
		b.WriteString(fmt.Sprintf("%sLastOnline: %s\n", indentationValues, friendPersistentInfo.LastOnline.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sLastOnline: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendPersistentInfo returns a new FriendPersistentInfo
func NewFriendPersistentInfo() *FriendPersistentInfo {
	return &FriendPersistentInfo{}
}
