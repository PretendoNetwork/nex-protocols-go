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
	Unknown  types.UInt32
	Unknown2 types.UInt32
	Unknown3 types.UInt32
	NEXToken types.String
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

	aei.Unknown.WriteTo(contentWritable)
	aei.Unknown2.WriteTo(contentWritable)
	aei.Unknown3.WriteTo(contentWritable)
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

	err = aei.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.Unknown. %s", err.Error())
	}

	err = aei.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.Unknown2. %s", err.Error())
	}

	err = aei.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.Unknown3. %s", err.Error())
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
	copied.Unknown = aei.Unknown.Copy().(types.UInt32)
	copied.Unknown2 = aei.Unknown2.Copy().(types.UInt32)
	copied.Unknown3 = aei.Unknown3.Copy().(types.UInt32)
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

	if !aei.Unknown.Equals(other.Unknown) {
		return false
	}

	if !aei.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !aei.Unknown3.Equals(other.Unknown3) {
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
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, aei.Unknown))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, aei.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, aei.Unknown3))
	b.WriteString(fmt.Sprintf("%sNEXToken: %s,\n", indentationValues, aei.NEXToken))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewAccountExtraInfo returns a new AccountExtraInfo
func NewAccountExtraInfo() AccountExtraInfo {
	return AccountExtraInfo{
		Unknown:  types.NewUInt32(0),
		Unknown2: types.NewUInt32(0),
		Unknown3: types.NewUInt32(0),
		NEXToken: types.NewString(""),
	}

}
