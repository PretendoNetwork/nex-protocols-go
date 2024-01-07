// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// BankTransactionParam holds data for the DataStore (Pokemon Bank) protocol
type BankTransactionParam struct {
	types.Structure
	DataID              *types.PrimitiveU64
	CurVersion          *types.PrimitiveU32
	UpdateVersion       *types.PrimitiveU32
	Size                *types.PrimitiveU32
	TransactionPassword *types.PrimitiveU64
}

// ExtractFrom extracts the BankTransactionParam from the given readable
func (bankTransactionParam *BankTransactionParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = bankTransactionParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read BankTransactionParam header. %s", err.Error())
	}

	err = bankTransactionParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.DataID from stream. %s", err.Error())
	}

	err = bankTransactionParam.CurVersion.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.CurVersion from stream. %s", err.Error())
	}

	err = bankTransactionParam.UpdateVersion.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.UpdateVersion from stream. %s", err.Error())
	}

	err = bankTransactionParam.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.Size from stream. %s", err.Error())
	}

	err = bankTransactionParam.TransactionPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.TransactionPassword from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the BankTransactionParam to the given writable
func (bankTransactionParam *BankTransactionParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	bankTransactionParam.DataID.WriteTo(contentWritable)
	bankTransactionParam.CurVersion.WriteTo(contentWritable)
	bankTransactionParam.UpdateVersion.WriteTo(contentWritable)
	bankTransactionParam.Size.WriteTo(contentWritable)
	bankTransactionParam.TransactionPassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	bankTransactionParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of BankTransactionParam
func (bankTransactionParam *BankTransactionParam) Copy() types.RVType {
	copied := NewBankTransactionParam()

	copied.StructureVersion = bankTransactionParam.StructureVersion

	copied.DataID = bankTransactionParam.DataID.Copy().(*types.PrimitiveU64)
	copied.CurVersion = bankTransactionParam.CurVersion.Copy().(*types.PrimitiveU32)
	copied.UpdateVersion = bankTransactionParam.UpdateVersion.Copy().(*types.PrimitiveU32)
	copied.Size = bankTransactionParam.Size.Copy().(*types.PrimitiveU32)
	copied.TransactionPassword = bankTransactionParam.TransactionPassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (bankTransactionParam *BankTransactionParam) Equals(o types.RVType) bool {
	if _, ok := o.(*BankTransactionParam); !ok {
		return false
	}

	other := o.(*BankTransactionParam)

	if bankTransactionParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !bankTransactionParam.DataID.Equals(other.DataID) {
		return false
	}

	if !bankTransactionParam.CurVersion.Equals(other.CurVersion) {
		return false
	}

	if !bankTransactionParam.UpdateVersion.Equals(other.UpdateVersion) {
		return false
	}

	if !bankTransactionParam.Size.Equals(other.Size) {
		return false
	}

	if !bankTransactionParam.TransactionPassword.Equals(other.TransactionPassword) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (bankTransactionParam *BankTransactionParam) String() string {
	return bankTransactionParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (bankTransactionParam *BankTransactionParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BankTransactionParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, bankTransactionParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, bankTransactionParam.DataID))
	b.WriteString(fmt.Sprintf("%sCurVersion: %s,\n", indentationValues, bankTransactionParam.CurVersion))
	b.WriteString(fmt.Sprintf("%sUpdateVersion: %s,\n", indentationValues, bankTransactionParam.UpdateVersion))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, bankTransactionParam.Size))
	b.WriteString(fmt.Sprintf("%sTransactionPassword: %s,\n", indentationValues, bankTransactionParam.TransactionPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBankTransactionParam returns a new BankTransactionParam
func NewBankTransactionParam() *BankTransactionParam {
	return &BankTransactionParam{
		DataID: types.NewPrimitiveU64(0),
		CurVersion: types.NewPrimitiveU32(0),
		UpdateVersion: types.NewPrimitiveU32(0),
		Size: types.NewPrimitiveU32(0),
		TransactionPassword: types.NewPrimitiveU64(0),
	}
}
