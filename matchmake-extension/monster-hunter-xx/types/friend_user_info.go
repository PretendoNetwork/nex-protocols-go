// Package types implements all the types used by the Matchmake Extension (Monster Hunter XX) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// FriendUserInfo holds data for the Matchmake Extension (Monster Hunter XX) protocol
type FriendUserInfo struct {
	nex.Structure
	PID      uint64
	Name     string
	Presence uint32
}

// ExtractFromStream extracts a FriendUserInfo structure from a stream
func (friendUserInfo *FriendUserInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	friendUserInfo.PID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserInfo.PID from stream. %s", err.Error())
	}

	friendUserInfo.Name, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserInfo.Name from stream. %s", err.Error())
	}

	friendUserInfo.Presence, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserInfo.Presence from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the FriendUserInfo and returns a byte array
func (friendUserInfo *FriendUserInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(friendUserInfo.PID)
	stream.WriteString(friendUserInfo.Name)
	stream.WriteUInt32LE(friendUserInfo.Presence)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendUserInfo
func (friendUserInfo *FriendUserInfo) Copy() nex.StructureInterface {
	copied := NewFriendUserInfo()

	copied.SetStructureVersion(friendUserInfo.StructureVersion())

	copied.PID = friendUserInfo.PID
	copied.Name = friendUserInfo.Name
	copied.Presence = friendUserInfo.Presence

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendUserInfo *FriendUserInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendUserInfo)

	if friendUserInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if friendUserInfo.PID != other.PID {
		return false
	}

	if friendUserInfo.Name != other.Name {
		return false
	}

	if friendUserInfo.Presence != other.Presence {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (friendUserInfo *FriendUserInfo) String() string {
	return friendUserInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (friendUserInfo *FriendUserInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendUserInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, friendUserInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, friendUserInfo.PID))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, friendUserInfo.Name))
	b.WriteString(fmt.Sprintf("%sPresence: %d,\n", indentationValues, friendUserInfo.Presence))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendUserInfo returns a new FriendUserInfo
func NewFriendUserInfo() *FriendUserInfo {
	return &FriendUserInfo{}
}
