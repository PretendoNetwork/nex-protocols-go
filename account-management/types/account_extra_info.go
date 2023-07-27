// Package types implements all the types used by the Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// AccountExtraInfo contains data for creating a new NNID on the network
type AccountExtraInfo struct {
	nex.Structure
	Unknown  uint32
	Unknown2 uint32
	Unknown3 uint32
	NEXToken string
}

// ExtractFromStream extracts a AccountExtraInfo structure from a stream
func (accountExtraInfo *AccountExtraInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	accountExtraInfo.Unknown, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.Unknown. %s", err.Error())
	}

	accountExtraInfo.Unknown2, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.Unknown2. %s", err.Error())
	}

	accountExtraInfo.Unknown3, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.Unknown3. %s", err.Error())
	}

	accountExtraInfo.NEXToken, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.NEXToken. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of AccountExtraInfo
func (accountExtraInfo *AccountExtraInfo) Copy() nex.StructureInterface {
	copied := NewAccountExtraInfo()

	copied.Unknown = accountExtraInfo.Unknown
	copied.Unknown2 = accountExtraInfo.Unknown2
	copied.Unknown3 = accountExtraInfo.Unknown3
	copied.NEXToken = accountExtraInfo.NEXToken

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (accountExtraInfo *AccountExtraInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*AccountExtraInfo)

	if accountExtraInfo.Unknown != other.Unknown {
		return false
	}

	if accountExtraInfo.Unknown2 != other.Unknown2 {
		return false
	}

	if accountExtraInfo.Unknown3 != other.Unknown3 {
		return false
	}

	if accountExtraInfo.NEXToken != other.NEXToken {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, accountExtraInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUnknown: %d,\n", indentationValues, accountExtraInfo.Unknown))
	b.WriteString(fmt.Sprintf("%sUnknown2: %d,\n", indentationValues, accountExtraInfo.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %d,\n", indentationValues, accountExtraInfo.Unknown3))
	b.WriteString(fmt.Sprintf("%sNEXToken: %q\n", indentationValues, accountExtraInfo.NEXToken))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewAccountExtraInfo returns a new AccountExtraInfo
func NewAccountExtraInfo() *AccountExtraInfo {
	return &AccountExtraInfo{}
}
