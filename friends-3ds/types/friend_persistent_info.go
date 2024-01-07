// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FriendPersistentInfo contains user settings
type FriendPersistentInfo struct {
	types.Structure
	*types.Data
	PID              *types.PID
	Region           *types.PrimitiveU8
	Country          *types.PrimitiveU8
	Area             *types.PrimitiveU8
	Language         *types.PrimitiveU8
	Platform         *types.PrimitiveU8
	GameKey          *GameKey
	Message          string
	MessageUpdatedAt *types.DateTime
	MiiModifiedAt    *types.DateTime
	LastOnline       *types.DateTime
}

// WriteTo writes the FriendPersistentInfo to the given writable
func (friendPersistentInfo *FriendPersistentInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	friendPersistentInfo.PID.WriteTo(contentWritable)
	friendPersistentInfo.Region.WriteTo(contentWritable)
	friendPersistentInfo.Country.WriteTo(contentWritable)
	friendPersistentInfo.Area.WriteTo(contentWritable)
	friendPersistentInfo.Language.WriteTo(contentWritable)
	friendPersistentInfo.Platform.WriteTo(contentWritable)
	friendPersistentInfo.GameKey.WriteTo(contentWritable)
	friendPersistentInfo.Message.WriteTo(contentWritable)
	friendPersistentInfo.MessageUpdatedAt.WriteTo(contentWritable)
	friendPersistentInfo.MiiModifiedAt.WriteTo(contentWritable)
	friendPersistentInfo.LastOnline.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	friendPersistentInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of FriendPersistentInfo
func (friendPersistentInfo *FriendPersistentInfo) Copy() types.RVType {
	copied := NewFriendPersistentInfo()

	copied.StructureVersion = friendPersistentInfo.StructureVersion

	copied.Data = friendPersistentInfo.Data.Copy().(*types.Data)

	copied.PID = friendPersistentInfo.PID.Copy()
	copied.Region = friendPersistentInfo.Region
	copied.Country = friendPersistentInfo.Country
	copied.Area = friendPersistentInfo.Area
	copied.Language = friendPersistentInfo.Language
	copied.Platform = friendPersistentInfo.Platform
	copied.GameKey = friendPersistentInfo.GameKey.Copy().(*GameKey)
	copied.Message = friendPersistentInfo.Message
	copied.MessageUpdatedAt = friendPersistentInfo.MessageUpdatedAt.Copy()
	copied.MiiModifiedAt = friendPersistentInfo.MiiModifiedAt.Copy()
	copied.LastOnline = friendPersistentInfo.LastOnline.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendPersistentInfo *FriendPersistentInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendPersistentInfo); !ok {
		return false
	}

	other := o.(*FriendPersistentInfo)

	if friendPersistentInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !friendPersistentInfo.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !friendPersistentInfo.PID.Equals(other.PID) {
		return false
	}

	if !friendPersistentInfo.Region.Equals(other.Region) {
		return false
	}

	if !friendPersistentInfo.Country.Equals(other.Country) {
		return false
	}

	if !friendPersistentInfo.Area.Equals(other.Area) {
		return false
	}

	if !friendPersistentInfo.Language.Equals(other.Language) {
		return false
	}

	if !friendPersistentInfo.Platform.Equals(other.Platform) {
		return false
	}

	if !friendPersistentInfo.GameKey.Equals(other.GameKey) {
		return false
	}

	if !friendPersistentInfo.Message.Equals(other.Message) {
		return false
	}

	if !friendPersistentInfo.MessageUpdatedAt.Equals(other.MessageUpdatedAt) {
		return false
	}

	if !friendPersistentInfo.MiiModifiedAt.Equals(other.MiiModifiedAt) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, friendPersistentInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, friendPersistentInfo.PID.FormatToString(indentationLevel+1)))
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

	if friendPersistentInfo.MiiModifiedAt != nil {
		b.WriteString(fmt.Sprintf("%sMiiModifiedAt: %s,\n", indentationValues, friendPersistentInfo.MiiModifiedAt.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sMiiModifiedAt: nil,\n", indentationValues))
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
