// Package types implements all the types used by the DataStore (Pokemon Bank) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// BankMigrationInfo holds data for the DataStore (Pokemon Bank) protocol
type BankMigrationInfo struct {
	types.Structure
	MigrationStatus *types.PrimitiveU32
	UpdatedTime     *types.DateTime
}

// ExtractFrom extracts the BankMigrationInfo from the given readable
func (bankMigrationInfo *BankMigrationInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = bankMigrationInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read BankMigrationInfo header. %s", err.Error())
	}

	err = bankMigrationInfo.MigrationStatus.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankMigrationInfo.MigrationStatus from stream. %s", err.Error())
	}

	err = bankMigrationInfo.UpdatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BankMigrationInfo.UpdatedTime from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the BankMigrationInfo to the given writable
func (bankMigrationInfo *BankMigrationInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	bankMigrationInfo.MigrationStatus.WriteTo(contentWritable)
	bankMigrationInfo.UpdatedTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	bankMigrationInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of BankMigrationInfo
func (bankMigrationInfo *BankMigrationInfo) Copy() types.RVType {
	copied := NewBankMigrationInfo()

	copied.StructureVersion = bankMigrationInfo.StructureVersion

	copied.MigrationStatus = bankMigrationInfo.MigrationStatus.Copy().(*types.PrimitiveU32)
	copied.UpdatedTime = bankMigrationInfo.UpdatedTime.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (bankMigrationInfo *BankMigrationInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*BankMigrationInfo); !ok {
		return false
	}

	other := o.(*BankMigrationInfo)

	if bankMigrationInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !bankMigrationInfo.MigrationStatus.Equals(other.MigrationStatus) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, bankMigrationInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sMigrationStatus: %s,\n", indentationValues, bankMigrationInfo.MigrationStatus))
	b.WriteString(fmt.Sprintf("%sUpdatedTime: %s\n", indentationValues, bankMigrationInfo.UpdatedTime.FormatToString(indentationLevel+1)))

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBankMigrationInfo returns a new BankMigrationInfo
func NewBankMigrationInfo() *BankMigrationInfo {
	return &BankMigrationInfo{
		MigrationStatus: types.NewPrimitiveU32(0),
		UpdatedTime: types.NewDateTime(0),
	}
}
