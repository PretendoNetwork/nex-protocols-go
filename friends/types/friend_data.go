// Package types implements all the types used by the Friends protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FriendData is a type within the Friends protocol
type FriendData struct {
	types.Structure
	PID            types.PID
	StrName        types.String
	ByRelationship types.UInt8
	UIDetails      types.UInt32
	StrStatus      types.String
}

// WriteTo writes the FriendData to the given writable
func (fd FriendData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	fd.PID.WriteTo(contentWritable)
	fd.StrName.WriteTo(contentWritable)
	fd.ByRelationship.WriteTo(contentWritable)
	fd.UIDetails.WriteTo(contentWritable)
	fd.StrStatus.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FriendData from the given readable
func (fd *FriendData) ExtractFrom(readable types.Readable) error {
	var err error

	err = fd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData header. %s", err.Error())
	}

	err = fd.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.PID. %s", err.Error())
	}

	err = fd.StrName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.StrName. %s", err.Error())
	}

	err = fd.ByRelationship.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.ByRelationship. %s", err.Error())
	}

	err = fd.UIDetails.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.UIDetails. %s", err.Error())
	}

	err = fd.StrStatus.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendData.StrStatus. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FriendData
func (fd FriendData) Copy() types.RVType {
	copied := NewFriendData()

	copied.StructureVersion = fd.StructureVersion
	copied.PID = fd.PID.Copy().(types.PID)
	copied.StrName = fd.StrName.Copy().(types.String)
	copied.ByRelationship = fd.ByRelationship.Copy().(types.UInt8)
	copied.UIDetails = fd.UIDetails.Copy().(types.UInt32)
	copied.StrStatus = fd.StrStatus.Copy().(types.String)

	return copied
}

// Equals checks if the given FriendData contains the same data as the current FriendData
func (fd FriendData) Equals(o types.RVType) bool {
	if _, ok := o.(FriendData); !ok {
		return false
	}

	other := o.(FriendData)

	if fd.StructureVersion != other.StructureVersion {
		return false
	}

	if !fd.PID.Equals(other.PID) {
		return false
	}

	if !fd.StrName.Equals(other.StrName) {
		return false
	}

	if !fd.ByRelationship.Equals(other.ByRelationship) {
		return false
	}

	if !fd.UIDetails.Equals(other.UIDetails) {
		return false
	}

	return fd.StrStatus.Equals(other.StrStatus)
}

// CopyRef copies the current value of the FriendData
// and returns a pointer to the new copy
func (fd FriendData) CopyRef() types.RVTypePtr {
	copied := fd.Copy().(FriendData)
	return &copied
}

// Deref takes a pointer to the FriendData
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (fd *FriendData) Deref() types.RVType {
	return *fd
}

// String returns the string representation of the FriendData
func (fd FriendData) String() string {
	return fd.FormatToString(0)
}

// FormatToString pretty-prints the FriendData using the provided indentation level
func (fd FriendData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendData{\n")
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, fd.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStrName: %s,\n", indentationValues, fd.StrName))
	b.WriteString(fmt.Sprintf("%sByRelationship: %s,\n", indentationValues, fd.ByRelationship))
	b.WriteString(fmt.Sprintf("%sUIDetails: %s,\n", indentationValues, fd.UIDetails))
	b.WriteString(fmt.Sprintf("%sStrStatus: %s,\n", indentationValues, fd.StrStatus))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendData returns a new FriendData
func NewFriendData() FriendData {
	return FriendData{
		PID:            types.NewPID(0),
		StrName:        types.NewString(""),
		ByRelationship: types.NewUInt8(0),
		UIDetails:      types.NewUInt32(0),
		StrStatus:      types.NewString(""),
	}

}
