// Package types implements all the types used by the AccountManagement protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// AccountExtraInfo is a type within the AccountManagement protocol
type AccountExtraInfo struct {
	types.Structure
	LocalFriendCode types.UInt64
	MoveCount       types.UInt32
	NEXToken        types.String
}

// ObjectID returns the object identifier of the type
func (aei AccountExtraInfo) ObjectID() types.RVType {
	return aei.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (aei AccountExtraInfo) DataObjectID() types.RVType {
	return types.NewString("AccountExtraInfo")
}

// WriteTo writes the AccountExtraInfo to the given writable
func (aei AccountExtraInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	aei.LocalFriendCode.WriteTo(contentWritable)
	aei.MoveCount.WriteTo(contentWritable)
	aei.NEXToken.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	aei.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the AccountExtraInfo from the given readable
func (aei *AccountExtraInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = aei.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo header. %s", err.Error())
	}

	err = aei.LocalFriendCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.LocalFriendCode. %s", err.Error())
	}

	err = aei.MoveCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.MoveCount. %s", err.Error())
	}

	err = aei.NEXToken.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.NEXToken. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of AccountExtraInfo
func (aei AccountExtraInfo) Copy() types.RVType {
	copied := NewAccountExtraInfo()

	copied.StructureVersion = aei.StructureVersion
	copied.LocalFriendCode = aei.LocalFriendCode.Copy().(types.UInt64)
	copied.MoveCount = aei.MoveCount.Copy().(types.UInt32)
	copied.NEXToken = aei.NEXToken.Copy().(types.String)

	return copied
}

// Equals checks if the given AccountExtraInfo contains the same data as the current AccountExtraInfo
func (aei AccountExtraInfo) Equals(o types.RVType) bool {
	if _, ok := o.(AccountExtraInfo); !ok {
		return false
	}

	other := o.(AccountExtraInfo)

	if aei.StructureVersion != other.StructureVersion {
		return false
	}

	if !aei.LocalFriendCode.Equals(other.LocalFriendCode) {
		return false
	}

	if !aei.MoveCount.Equals(other.MoveCount) {
		return false
	}

	return aei.NEXToken.Equals(other.NEXToken)
}

// CopyRef copies the current value of the AccountExtraInfo
// and returns a pointer to the new copy
func (aei AccountExtraInfo) CopyRef() types.RVTypePtr {
	copied := aei.Copy().(AccountExtraInfo)
	return &copied
}

// Deref takes a pointer to the AccountExtraInfo
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (aei *AccountExtraInfo) Deref() types.RVType {
	return *aei
}

// String returns the string representation of the AccountExtraInfo
func (aei AccountExtraInfo) String() string {
	return aei.FormatToString(0)
}

// FormatToString pretty-prints the AccountExtraInfo using the provided indentation level
func (aei AccountExtraInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("AccountExtraInfo{\n")
	b.WriteString(fmt.Sprintf("%sLocalFriendCode: %s,\n", indentationValues, aei.LocalFriendCode))
	b.WriteString(fmt.Sprintf("%sMoveCount: %s,\n", indentationValues, aei.MoveCount))
	b.WriteString(fmt.Sprintf("%sNEXToken: %s,\n", indentationValues, aei.NEXToken))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewAccountExtraInfo returns a new AccountExtraInfo
func NewAccountExtraInfo() AccountExtraInfo {
	return AccountExtraInfo{
		LocalFriendCode: types.NewUInt64(0),
		MoveCount:       types.NewUInt32(0),
		NEXToken:        types.NewString(""),
	}

}
