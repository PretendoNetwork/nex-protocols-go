// Package types implements all the types used by the Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// AccountData contains data for creating a new NNID on the network
type AccountData struct {
	types.Structure
	PID                *types.PID
	StrName            *types.String
	UIGroups           *types.PrimitiveU32
	StrEmail           *types.String
	DTCreationDate     *types.DateTime
	DTEffectiveDate    *types.DateTime
	StrNotEffectiveMsg *types.String
	DTExpiryDate       *types.DateTime
	StrExpiredMsg      *types.String
}

// WriteTo writes the AccountData to the given writable
func (accountData *AccountData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	accountData.PID.WriteTo(contentWritable)
	accountData.StrName.WriteTo(contentWritable)
	accountData.UIGroups.WriteTo(contentWritable)
	accountData.StrEmail.WriteTo(contentWritable)
	accountData.DTCreationDate.WriteTo(contentWritable)
	accountData.DTEffectiveDate.WriteTo(contentWritable)
	accountData.StrNotEffectiveMsg.WriteTo(contentWritable)
	accountData.DTExpiryDate.WriteTo(contentWritable)
	accountData.StrExpiredMsg.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	accountData.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the AccountData from the given readable
func (accountData *AccountData) ExtractFrom(readable types.Readable) error {
	var err error

	if err = accountData.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read AccountData header. %s", err.Error())
	}

	err = accountData.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.PID. %s", err.Error())
	}

	err = accountData.StrName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.StrName. %s", err.Error())
	}

	err = accountData.UIGroups.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.UIGroups. %s", err.Error())
	}

	err = accountData.StrEmail.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.StrEmail. %s", err.Error())
	}

	err = accountData.DTCreationDate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.DTCreationDate. %s", err.Error())
	}

	err = accountData.DTEffectiveDate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.DTEffectiveDate. %s", err.Error())
	}

	err = accountData.StrNotEffectiveMsg.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.StrNotEffectiveMsg. %s", err.Error())
	}

	err = accountData.DTExpiryDate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.DTExpiryDate. %s", err.Error())
	}

	err = accountData.StrExpiredMsg.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.StrExpiredMsg. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of AccountData
func (accountData *AccountData) Copy() types.RVType {
	copied := NewAccountData()

	copied.StructureVersion = accountData.StructureVersion

	copied.PID = accountData.PID.Copy().(*types.PID)
	copied.StrName = accountData.StrName.Copy().(*types.String)
	copied.UIGroups = accountData.UIGroups.Copy().(*types.PrimitiveU32)
	copied.StrEmail = accountData.StrEmail.Copy().(*types.String)

	copied.DTCreationDate = accountData.DTCreationDate.Copy().Copy().(*types.DateTime)

	copied.DTEffectiveDate = accountData.DTEffectiveDate.Copy().Copy().(*types.DateTime)

	copied.StrNotEffectiveMsg = accountData.StrNotEffectiveMsg.Copy().(*types.String)

	copied.DTExpiryDate = accountData.DTExpiryDate.Copy().(*types.DateTime)

	copied.StrExpiredMsg = accountData.StrExpiredMsg.Copy().(*types.String)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (accountData *AccountData) Equals(o types.RVType) bool {
	if _, ok := o.(*AccountData); !ok {
		return false
	}

	other := o.(*AccountData)

	if accountData.StructureVersion != other.StructureVersion {
		return false
	}

	if !accountData.PID.Equals(other.PID) {
		return false
	}

	if !accountData.StrName.Equals(other.StrName) {
		return false
	}

	if !accountData.UIGroups.Equals(other.UIGroups) {
		return false
	}

	if !accountData.StrEmail.Equals(other.StrEmail) {
		return false
	}

	if !accountData.DTCreationDate.Equals(other.DTCreationDate) {
		return false
	}

	if !accountData.DTEffectiveDate.Equals(other.DTEffectiveDate) {
		return false
	}

	if !accountData.StrNotEffectiveMsg.Equals(other.StrNotEffectiveMsg) {
		return false
	}

	if !accountData.DTExpiryDate.Equals(other.DTExpiryDate) {
		return false
	}

	if !accountData.StrExpiredMsg.Equals(other.StrExpiredMsg) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (accountData *AccountData) String() string {
	return accountData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (accountData *AccountData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("AccountData{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, accountData.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, accountData.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStrName: %s,\n", indentationValues, accountData.StrName))
	b.WriteString(fmt.Sprintf("%sUIGroups: %s,\n", indentationValues, accountData.UIGroups))
	b.WriteString(fmt.Sprintf("%sStrEmail: %s,\n", indentationValues, accountData.StrEmail))
	b.WriteString(fmt.Sprintf("%sDTCreationDate: %s,\n", indentationValues, accountData.DTCreationDate.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDTEffectiveDate: %s,\n", indentationValues, accountData.DTEffectiveDate.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStrNotEffectiveMsg: %s,\n", indentationValues, accountData.StrNotEffectiveMsg))
	b.WriteString(fmt.Sprintf("%sDTExpiryDate: %s,\n", indentationValues, accountData.DTExpiryDate.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStrExpiredMsg: %s\n", indentationValues, accountData.StrExpiredMsg))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewAccountData returns a new AccountData
func NewAccountData() *AccountData {
	return &AccountData{
		PID: types.NewPID(0),
		StrName: types.NewString(""),
		UIGroups: types.NewPrimitiveU32(0),
		StrEmail: types.NewString(""),
		DTCreationDate: types.NewDateTime(0),
		DTEffectiveDate: types.NewDateTime(0),
		StrNotEffectiveMsg: types.NewString(""),
		DTExpiryDate: types.NewDateTime(0),
		StrExpiredMsg: types.NewString(""),
	}
}
