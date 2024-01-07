// Package types implements all the types used by the Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// AccountExtraInfo contains data for creating a new NNID on the network
type AccountExtraInfo struct {
	types.Structure
	Unknown  *types.PrimitiveU32
	Unknown2 *types.PrimitiveU32
	Unknown3 *types.PrimitiveU32
	NEXToken *types.String
}

// WriteTo writes the AccountExtraInfo to the given writable
func (accountExtraInfo *AccountExtraInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	accountExtraInfo.Unknown.WriteTo(contentWritable)
	accountExtraInfo.Unknown2.WriteTo(contentWritable)
	accountExtraInfo.Unknown3.WriteTo(contentWritable)
	accountExtraInfo.NEXToken.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	accountExtraInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the AccountExtraInfo from the given readable
func (accountExtraInfo *AccountExtraInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = accountExtraInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read AccountExtraInfo header. %s", err.Error())
	}

	err = accountExtraInfo.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.Unknown. %s", err.Error())
	}

	err = accountExtraInfo.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.Unknown2. %s", err.Error())
	}

	err = accountExtraInfo.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.Unknown3. %s", err.Error())
	}

	err = accountExtraInfo.NEXToken.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.NEXToken. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of AccountExtraInfo
func (accountExtraInfo *AccountExtraInfo) Copy() types.RVType {
	copied := NewAccountExtraInfo()

	copied.StructureVersion = accountExtraInfo.StructureVersion

	copied.Unknown = accountExtraInfo.Unknown.Copy().(*types.PrimitiveU32)
	copied.Unknown2 = accountExtraInfo.Unknown2.Copy().(*types.PrimitiveU32)
	copied.Unknown3 = accountExtraInfo.Unknown3.Copy().(*types.PrimitiveU32)
	copied.NEXToken = accountExtraInfo.NEXToken.Copy().(*types.String)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (accountExtraInfo *AccountExtraInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*AccountExtraInfo); !ok {
		return false
	}

	other := o.(*AccountExtraInfo)

	if accountExtraInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !accountExtraInfo.Unknown.Equals(other.Unknown) {
		return false
	}

	if !accountExtraInfo.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !accountExtraInfo.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !accountExtraInfo.NEXToken.Equals(other.NEXToken) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (accountExtraInfo *AccountExtraInfo) String() string {
	return accountExtraInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (accountExtraInfo *AccountExtraInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("AccountExtraInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, accountExtraInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, accountExtraInfo.Unknown))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, accountExtraInfo.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, accountExtraInfo.Unknown3))
	b.WriteString(fmt.Sprintf("%sNEXToken: %s\n", indentationValues, accountExtraInfo.NEXToken))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewAccountExtraInfo returns a new AccountExtraInfo
func NewAccountExtraInfo() *AccountExtraInfo {
	return &AccountExtraInfo{
		Unknown: types.NewPrimitiveU32(0),
		Unknown2: types.NewPrimitiveU32(0),
		Unknown3: types.NewPrimitiveU32(0),
		NEXToken: types.NewString(""),
	}
}
