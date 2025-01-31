// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// BankTransactionParam is a type within the DataStore protocol
type BankTransactionParam struct {
	types.Structure
	DataID              types.UInt64
	CurVersion          types.UInt32
	UpdateVersion       types.UInt32
	Size                types.UInt32
	TransactionPassword types.UInt64
}

// WriteTo writes the BankTransactionParam to the given writable
func (btp BankTransactionParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	btp.DataID.WriteTo(contentWritable)
	btp.CurVersion.WriteTo(contentWritable)
	btp.UpdateVersion.WriteTo(contentWritable)
	btp.Size.WriteTo(contentWritable)
	btp.TransactionPassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	btp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the BankTransactionParam from the given readable
func (btp *BankTransactionParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = btp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam header. %s", err.Error())
	}

	err = btp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.DataID. %s", err.Error())
	}

	err = btp.CurVersion.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.CurVersion. %s", err.Error())
	}

	err = btp.UpdateVersion.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.UpdateVersion. %s", err.Error())
	}

	err = btp.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.Size. %s", err.Error())
	}

	err = btp.TransactionPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.TransactionPassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BankTransactionParam
func (btp BankTransactionParam) Copy() types.RVType {
	copied := NewBankTransactionParam()

	copied.StructureVersion = btp.StructureVersion
	copied.DataID = btp.DataID.Copy().(types.UInt64)
	copied.CurVersion = btp.CurVersion.Copy().(types.UInt32)
	copied.UpdateVersion = btp.UpdateVersion.Copy().(types.UInt32)
	copied.Size = btp.Size.Copy().(types.UInt32)
	copied.TransactionPassword = btp.TransactionPassword.Copy().(types.UInt64)

	return copied
}

// Equals checks if the given BankTransactionParam contains the same data as the current BankTransactionParam
func (btp BankTransactionParam) Equals(o types.RVType) bool {
	if _, ok := o.(BankTransactionParam); !ok {
		return false
	}

	other := o.(BankTransactionParam)

	if btp.StructureVersion != other.StructureVersion {
		return false
	}

	if !btp.DataID.Equals(other.DataID) {
		return false
	}

	if !btp.CurVersion.Equals(other.CurVersion) {
		return false
	}

	if !btp.UpdateVersion.Equals(other.UpdateVersion) {
		return false
	}

	if !btp.Size.Equals(other.Size) {
		return false
	}

	return btp.TransactionPassword.Equals(other.TransactionPassword)
}

// CopyRef copies the current value of the BankTransactionParam
// and returns a pointer to the new copy
func (btp BankTransactionParam) CopyRef() types.RVTypePtr {
	copied := btp.Copy().(BankTransactionParam)
	return &copied
}

// Deref takes a pointer to the BankTransactionParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (btp *BankTransactionParam) Deref() types.RVType {
	return *btp
}

// String returns the string representation of the BankTransactionParam
func (btp BankTransactionParam) String() string {
	return btp.FormatToString(0)
}

// FormatToString pretty-prints the BankTransactionParam using the provided indentation level
func (btp BankTransactionParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BankTransactionParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, btp.DataID))
	b.WriteString(fmt.Sprintf("%sCurVersion: %s,\n", indentationValues, btp.CurVersion))
	b.WriteString(fmt.Sprintf("%sUpdateVersion: %s,\n", indentationValues, btp.UpdateVersion))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, btp.Size))
	b.WriteString(fmt.Sprintf("%sTransactionPassword: %s,\n", indentationValues, btp.TransactionPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBankTransactionParam returns a new BankTransactionParam
func NewBankTransactionParam() BankTransactionParam {
	return BankTransactionParam{
		DataID:              types.NewUInt64(0),
		CurVersion:          types.NewUInt32(0),
		UpdateVersion:       types.NewUInt32(0),
		Size:                types.NewUInt32(0),
		TransactionPassword: types.NewUInt64(0),
	}

}
