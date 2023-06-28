package account_management_types

import (
	"fmt"

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

// NewAccountExtraInfo returns a new AccountExtraInfo
func NewAccountExtraInfo() *AccountExtraInfo {
	return &AccountExtraInfo{}
}
