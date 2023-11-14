// Package types implements all the types used by the Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// AccountData contains data for creating a new NNID on the network
type AccountData struct {
	nex.Structure
	PID                *nex.PID
	StrName            string
	UIGroups           uint32
	strEmail           string
	DTCreationDate     *nex.DateTime
	DTEffectiveDate    *nex.DateTime
	StrNotEffectiveMsg string
	DTExpiryDate       *nex.DateTime
	StrExpiredMsg      string
}

// ExtractFromStream extracts a AccountData structure from a stream
func (accountData *AccountData) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	accountData.PID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.PID. %s", err.Error())
	}

	accountData.StrName, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.StrName. %s", err.Error())
	}

	accountData.UIGroups, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.UIGroups. %s", err.Error())
	}

	accountData.strEmail, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.strEmail. %s", err.Error())
	}

	accountData.DTCreationDate, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.DTCreationDate. %s", err.Error())
	}

	accountData.DTEffectiveDate, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.DTEffectiveDate. %s", err.Error())
	}

	accountData.StrNotEffectiveMsg, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.StrNotEffectiveMsg. %s", err.Error())
	}

	accountData.DTExpiryDate, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.DTExpiryDate. %s", err.Error())
	}

	accountData.StrExpiredMsg, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountData.StrExpiredMsg. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of AccountData
func (accountData *AccountData) Copy() nex.StructureInterface {
	copied := NewAccountData()

	copied.SetStructureVersion(accountData.StructureVersion())

	copied.PID = accountData.PID.Copy()
	copied.StrName = accountData.StrName
	copied.UIGroups = accountData.UIGroups
	copied.strEmail = accountData.strEmail

	if accountData.DTCreationDate != nil {
		copied.DTCreationDate = accountData.DTCreationDate.Copy()
	}

	if accountData.DTEffectiveDate != nil {
		copied.DTEffectiveDate = accountData.DTEffectiveDate.Copy()
	}

	copied.StrNotEffectiveMsg = accountData.StrNotEffectiveMsg

	if accountData.DTExpiryDate != nil {
		copied.DTExpiryDate = accountData.DTExpiryDate.Copy()
	}

	copied.StrExpiredMsg = accountData.StrExpiredMsg

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (accountData *AccountData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*AccountData)

	if accountData.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !accountData.PID.Equals(other.PID) {
		return false
	}

	if accountData.StrName != other.StrName {
		return false
	}

	if accountData.UIGroups != other.UIGroups {
		return false
	}

	if accountData.strEmail != other.strEmail {
		return false
	}

	if accountData.DTCreationDate != nil && other.DTCreationDate == nil {
		return false
	}

	if accountData.DTCreationDate == nil && other.DTCreationDate != nil {
		return false
	}

	if accountData.DTCreationDate != nil && other.DTCreationDate != nil {
		if !accountData.DTCreationDate.Equals(other.DTCreationDate) {
			return false
		}
	}

	if accountData.DTEffectiveDate != nil && other.DTEffectiveDate == nil {
		return false
	}

	if accountData.DTEffectiveDate == nil && other.DTEffectiveDate != nil {
		return false
	}

	if accountData.DTEffectiveDate != nil && other.DTEffectiveDate != nil {
		if !accountData.DTEffectiveDate.Equals(other.DTEffectiveDate) {
			return false
		}
	}

	if accountData.StrNotEffectiveMsg != other.StrNotEffectiveMsg {
		return false
	}

	if accountData.DTExpiryDate != nil && other.DTExpiryDate == nil {
		return false
	}

	if accountData.DTExpiryDate == nil && other.DTExpiryDate != nil {
		return false
	}

	if accountData.DTExpiryDate != nil && other.DTExpiryDate != nil {
		if !accountData.DTExpiryDate.Equals(other.DTExpiryDate) {
			return false
		}
	}

	if accountData.StrExpiredMsg != other.StrExpiredMsg {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, accountData.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, accountData.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStrName: %q,\n", indentationValues, accountData.StrName))
	b.WriteString(fmt.Sprintf("%sUIGroups: %d,\n", indentationValues, accountData.UIGroups))
	b.WriteString(fmt.Sprintf("%sstrEmail: %q,\n", indentationValues, accountData.strEmail))

	if accountData.DTCreationDate != nil {
		b.WriteString(fmt.Sprintf("%sDTCreationDate: %s,\n", indentationValues, accountData.DTCreationDate.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sDTCreationDate: nil,\n", indentationValues))

	}

	if accountData.DTEffectiveDate != nil {
		b.WriteString(fmt.Sprintf("%sDTEffectiveDate: %s,\n", indentationValues, accountData.DTEffectiveDate.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sDTEffectiveDate: nil,\n", indentationValues))

	}

	b.WriteString(fmt.Sprintf("%sStrNotEffectiveMsg: %q,\n", indentationValues, accountData.StrNotEffectiveMsg))

	if accountData.DTExpiryDate != nil {
		b.WriteString(fmt.Sprintf("%sDTExpiryDate: %s,\n", indentationValues, accountData.DTExpiryDate.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sDTExpiryDate: nil,\n", indentationValues))

	}

	b.WriteString(fmt.Sprintf("%sStrExpiredMsg: %q\n", indentationValues, accountData.StrExpiredMsg))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewAccountData returns a new AccountData
func NewAccountData() *AccountData {
	return &AccountData{}
}
