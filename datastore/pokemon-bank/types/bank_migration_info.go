// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// BankMigrationInfo holds data for the DataStore (Pokemon Bank) protocol
type BankMigrationInfo struct {
	nex.Structure
	MigrationStatus uint32
	UpdatedTime     *nex.DateTime
}

// ExtractFromStream extracts a BankMigrationInfo structure from a stream
func (bankMigrationInfo *BankMigrationInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	bankMigrationInfo.MigrationStatus, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract BankMigrationInfo.MigrationStatus from stream. %s", err.Error())
	}

	bankMigrationInfo.UpdatedTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract BankMigrationInfo.UpdatedTime from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the BankMigrationInfo and returns a byte array
func (bankMigrationInfo *BankMigrationInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(bankMigrationInfo.MigrationStatus)
	stream.WriteDateTime(bankMigrationInfo.UpdatedTime)

	return stream.Bytes()
}

// Copy returns a new copied instance of BankMigrationInfo
func (bankMigrationInfo *BankMigrationInfo) Copy() nex.StructureInterface {
	copied := NewBankMigrationInfo()

	copied.MigrationStatus = bankMigrationInfo.MigrationStatus
	copied.UpdatedTime = bankMigrationInfo.UpdatedTime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (bankMigrationInfo *BankMigrationInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*BankMigrationInfo)

	if bankMigrationInfo.MigrationStatus != other.MigrationStatus {
		return false
	}

	if !bankMigrationInfo.UpdatedTime.Equals(other.UpdatedTime) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (bankMigrationInfo *BankMigrationInfo) String() string {
	return bankMigrationInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (bankMigrationInfo *BankMigrationInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BankMigrationInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, bankMigrationInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sMigrationStatus: %d,\n", indentationValues, bankMigrationInfo.MigrationStatus))

	if bankMigrationInfo.UpdatedTime != nil {
		b.WriteString(fmt.Sprintf("%sUpdatedTime: %s\n", indentationValues, bankMigrationInfo.UpdatedTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUpdatedTime: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBankMigrationInfo returns a new BankMigrationInfo
func NewBankMigrationInfo() *BankMigrationInfo {
	return &BankMigrationInfo{}
}
