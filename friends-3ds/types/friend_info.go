// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FriendInfo is a data structure used by the Friends 3DS protocol to hold information about a friends Mii
type FriendInfo struct {
	types.Structure
	PID     *types.PID
	Unknown *types.DateTime
}

// WriteTo writes the FriendInfo to the given writable
func (friendInfo *FriendInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	friendInfo.PID.WriteTo(contentWritable)
	friendInfo.Unknown.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	friendInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendInfo from the given readable
func (friendInfo *FriendInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = friendInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read FriendInfo header. %s", err.Error())
	}

	err = friendInfo.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo.PID. %s", err.Error())
	}

	err = friendInfo.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo.Unknown. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendInfo
func (friendInfo *FriendInfo) Copy() types.RVType {
	copied := NewFriendInfo()

	copied.StructureVersion = friendInfo.StructureVersion

	copied.PID = friendInfo.PID.Copy()

	copied.Unknown = friendInfo.Unknown.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendInfo *FriendInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendInfo); !ok {
		return false
	}

	other := o.(*FriendInfo)

	if friendInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !friendInfo.PID.Equals(other.PID) {
		return false
	}

	if !friendInfo.Unknown.Equals(other.Unknown) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, friendInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, friendInfo.PID.FormatToString(indentationLevel+1)))

	if friendInfo.Unknown != nil {
		b.WriteString(fmt.Sprintf("%sUnknown: %s\n", indentationValues, friendInfo.Unknown.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUnknown: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendInfo returns a new FriendInfo
func NewFriendInfo() *FriendInfo {
	return &FriendInfo{}
}
