// Package types implements all the types used by the Matchmake Extension (Monster Hunter XX) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FriendUserInfo holds data for the Matchmake Extension (Monster Hunter XX) protocol
type FriendUserInfo struct {
	types.Structure
	PID      *types.PID
	Name     string
	Presence *types.PrimitiveU32
}

// ExtractFrom extracts the FriendUserInfo from the given readable
func (friendUserInfo *FriendUserInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = friendUserInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read FriendUserInfo header. %s", err.Error())
	}

	err = friendUserInfo.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserInfo.PID from stream. %s", err.Error())
	}

	err = friendUserInfo.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserInfo.Name from stream. %s", err.Error())
	}

	err = friendUserInfo.Presence.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserInfo.Presence from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the FriendUserInfo to the given writable
func (friendUserInfo *FriendUserInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	friendUserInfo.PID.WriteTo(contentWritable)
	friendUserInfo.Name.WriteTo(contentWritable)
	friendUserInfo.Presence.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	friendUserInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of FriendUserInfo
func (friendUserInfo *FriendUserInfo) Copy() types.RVType {
	copied := NewFriendUserInfo()

	copied.StructureVersion = friendUserInfo.StructureVersion

	copied.PID = friendUserInfo.PID.Copy()
	copied.Name = friendUserInfo.Name
	copied.Presence = friendUserInfo.Presence

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendUserInfo *FriendUserInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendUserInfo); !ok {
		return false
	}

	other := o.(*FriendUserInfo)

	if friendUserInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if friendUserInfo.PID.Equals(other.PID) {
		return false
	}

	if !friendUserInfo.Name.Equals(other.Name) {
		return false
	}

	if !friendUserInfo.Presence.Equals(other.Presence) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, friendUserInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, friendUserInfo.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, friendUserInfo.Name))
	b.WriteString(fmt.Sprintf("%sPresence: %d,\n", indentationValues, friendUserInfo.Presence))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendUserInfo returns a new FriendUserInfo
func NewFriendUserInfo() *FriendUserInfo {
	return &FriendUserInfo{}
}
