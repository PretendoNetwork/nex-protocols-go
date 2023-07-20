// Package friends_types implements all the types used by the Friends protocol
package friends_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// FriendData contains data relating to a friend
type FriendData struct {
	nex.Structure
	PID            uint32
	StrName        string
	ByRelationship uint8
	UIDetails      uint32
	StrStatus      string
}

// ExtractFromStream extracts a FriendData structure from a stream
func (friendData *FriendData) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	friendData.PID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.PID. %s", err.Error())
	}

	friendData.StrName, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.StrName. %s", err.Error())
	}

	friendData.ByRelationship, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.ByRelationship. %s", err.Error())
	}

	friendData.UIDetails, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.UIDetails. %s", err.Error())
	}

	friendData.StrStatus, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.StrStatus. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendData
func (friendData *FriendData) Copy() nex.StructureInterface {
	copied := NewFriendData()

	copied.PID = friendData.PID
	copied.StrName = friendData.StrName
	copied.ByRelationship = friendData.ByRelationship
	copied.UIDetails = friendData.UIDetails
	copied.StrStatus = friendData.StrStatus

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendData *FriendData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendData)

	if friendData.PID != other.PID {
		return false
	}

	if friendData.StrName != other.StrName {
		return false
	}

	if friendData.ByRelationship != other.ByRelationship {
		return false
	}

	if friendData.UIDetails != other.UIDetails {
		return false
	}

	if friendData.StrStatus != other.StrStatus {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (friendData *FriendData) String() string {
	return friendData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (friendData *FriendData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendData{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, friendData.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, friendData.PID))
	b.WriteString(fmt.Sprintf("%sStrName: %q,\n", indentationValues, friendData.StrName))
	b.WriteString(fmt.Sprintf("%sByRelationship: %d,\n", indentationValues, friendData.ByRelationship))
	b.WriteString(fmt.Sprintf("%sUIDetails: %d,\n", indentationValues, friendData.UIDetails))
	b.WriteString(fmt.Sprintf("%sStrStatus: %q\n", indentationValues, friendData.StrStatus))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendData returns a new FriendData
func NewFriendData() *FriendData {
	return &FriendData{}
}
