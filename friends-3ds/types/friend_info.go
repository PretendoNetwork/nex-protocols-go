// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// FriendInfo is a data structure used by the Friends 3DS protocol to hold information about a friends Mii
type FriendInfo struct {
	nex.Structure
	PID     *nex.PID
	Unknown *nex.DateTime
}

// Bytes encodes the FriendInfo and returns a byte array
func (friendInfo *FriendInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WritePID(friendInfo.PID)
	stream.WriteDateTime(friendInfo.Unknown)

	return stream.Bytes()
}

func (friendInfo *FriendInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	friendInfo.PID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo.PID. %s", err.Error())
	}

	friendInfo.Unknown, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract FriendInfo.PID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendInfo
func (friendInfo *FriendInfo) Copy() nex.StructureInterface {
	copied := NewFriendInfo()

	copied.SetStructureVersion(friendInfo.StructureVersion())

	if friendInfo.PID != nil {
		copied.PID = friendInfo.PID.Copy()
	}

	if friendInfo.Unknown != nil {
		copied.Unknown = friendInfo.Unknown.Copy()
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendInfo *FriendInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendInfo)

	if friendInfo.StructureVersion() != other.StructureVersion() {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, friendInfo.StructureVersion()))
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
