// Package types implements all the types used by the Friends protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FriendData contains data relating to a friend
type FriendData struct {
	types.Structure
	PID            *types.PID
	StrName        string
	ByRelationship *types.PrimitiveU8
	UIDetails      *types.PrimitiveU32
	StrStatus      string
}

// ExtractFrom extracts the FriendData from the given readable
func (friendData *FriendData) ExtractFrom(readable types.Readable) error {
	var err error

	if err = friendData.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read FriendData header. %s", err.Error())
	}

	err = friendData.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.PID. %s", err.Error())
	}

	err = friendData.StrName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.StrName. %s", err.Error())
	}

	err = friendData.ByRelationship.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.ByRelationship. %s", err.Error())
	}

	err = friendData.UIDetails.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.UIDetails. %s", err.Error())
	}

	err = friendData.StrStatus.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.StrStatus. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendData
func (friendData *FriendData) Copy() types.RVType {
	copied := NewFriendData()

	copied.StructureVersion = friendData.StructureVersion

	copied.PID = friendData.PID.Copy()
	copied.StrName = friendData.StrName
	copied.ByRelationship = friendData.ByRelationship
	copied.UIDetails = friendData.UIDetails
	copied.StrStatus = friendData.StrStatus

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendData *FriendData) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendData); !ok {
		return false
	}

	other := o.(*FriendData)

	if friendData.StructureVersion != other.StructureVersion {
		return false
	}

	if !friendData.PID.Equals(other.PID) {
		return false
	}

	if !friendData.StrName.Equals(other.StrName) {
		return false
	}

	if !friendData.ByRelationship.Equals(other.ByRelationship) {
		return false
	}

	if !friendData.UIDetails.Equals(other.UIDetails) {
		return false
	}

	if !friendData.StrStatus.Equals(other.StrStatus) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, friendData.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, friendData.PID.FormatToString(indentationLevel+1)))
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
