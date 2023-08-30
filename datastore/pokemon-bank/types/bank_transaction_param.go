// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// BankTransactionParam holds data for the DataStore (Pokemon Bank) protocol
type BankTransactionParam struct {
	nex.Structure
	DataID              uint64
	CurVersion          uint32
	UpdateVersion       uint32
	Size                uint32
	TransactionPassword uint64
}

// ExtractFromStream extracts a BankTransactionParam structure from a stream
func (bankTransactionParam *BankTransactionParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	bankTransactionParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.DataID from stream. %s", err.Error())
	}

	bankTransactionParam.CurVersion, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.CurVersion from stream. %s", err.Error())
	}

	bankTransactionParam.UpdateVersion, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.UpdateVersion from stream. %s", err.Error())
	}

	bankTransactionParam.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.Size from stream. %s", err.Error())
	}

	bankTransactionParam.TransactionPassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract BankTransactionParam.TransactionPassword from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the BankTransactionParam and returns a byte array
func (bankTransactionParam *BankTransactionParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(bankTransactionParam.DataID)
	stream.WriteUInt32LE(bankTransactionParam.CurVersion)
	stream.WriteUInt32LE(bankTransactionParam.UpdateVersion)
	stream.WriteUInt32LE(bankTransactionParam.Size)
	stream.WriteUInt64LE(bankTransactionParam.TransactionPassword)

	return stream.Bytes()
}

// Copy returns a new copied instance of BankTransactionParam
func (bankTransactionParam *BankTransactionParam) Copy() nex.StructureInterface {
	copied := NewBankTransactionParam()

	copied.SetStructureVersion(bankTransactionParam.StructureVersion())

	copied.DataID = bankTransactionParam.DataID
	copied.CurVersion = bankTransactionParam.CurVersion
	copied.UpdateVersion = bankTransactionParam.UpdateVersion
	copied.Size = bankTransactionParam.Size
	copied.TransactionPassword = bankTransactionParam.TransactionPassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (bankTransactionParam *BankTransactionParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*BankTransactionParam)

	if bankTransactionParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if bankTransactionParam.DataID != other.DataID {
		return false
	}

	if bankTransactionParam.CurVersion != other.CurVersion {
		return false
	}

	if bankTransactionParam.UpdateVersion != other.UpdateVersion {
		return false
	}

	if bankTransactionParam.Size != other.Size {
		return false
	}

	if bankTransactionParam.TransactionPassword != other.TransactionPassword {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, bankTransactionParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, bankTransactionParam.DataID))
	b.WriteString(fmt.Sprintf("%sCurVersion: %d,\n", indentationValues, bankTransactionParam.CurVersion))
	b.WriteString(fmt.Sprintf("%sUpdateVersion: %d,\n", indentationValues, bankTransactionParam.UpdateVersion))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, bankTransactionParam.Size))
	b.WriteString(fmt.Sprintf("%sTransactionPassword: %d,\n", indentationValues, bankTransactionParam.TransactionPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBankTransactionParam returns a new BankTransactionParam
func NewBankTransactionParam() *BankTransactionParam {
	return &BankTransactionParam{}
}
