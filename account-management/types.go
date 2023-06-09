package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu "github.com/PretendoNetwork/nex-protocols-go/friends/wiiu"
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

// NintendoCreateAccountData contains data for creating a new NNID on the network
type NintendoCreateAccountData struct {
	nex.Structure
	NNAInfo  *friends_wiiu.NNAInfo
	Token    string
	Birthday *nex.DateTime
	Unknown  uint64
}

// ExtractFromStream extracts a NintendoCreateAccountData structure from a stream
func (nintendoCreateAccountData *NintendoCreateAccountData) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	nnaInfo, err := stream.ReadStructure(friends_wiiu.NewNNAInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.NNAInfo from stream. %s", err.Error())
	}

	nintendoCreateAccountData.NNAInfo = nnaInfo.(*friends_wiiu.NNAInfo)
	nintendoCreateAccountData.Token, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.Token from stream. %s", err.Error())
	}

	nintendoCreateAccountData.Birthday, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.Birthday from stream. %s", err.Error())
	}

	nintendoCreateAccountData.Unknown, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.Unknown from stream. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoCreateAccountData
func (nintendoCreateAccountData *NintendoCreateAccountData) Copy() nex.StructureInterface {
	copied := NewNintendoCreateAccountData()

	copied.NNAInfo = nintendoCreateAccountData.NNAInfo.Copy().(*friends_wiiu.NNAInfo)
	copied.Token = nintendoCreateAccountData.Token
	copied.Birthday = nintendoCreateAccountData.Birthday.Copy()
	copied.Unknown = nintendoCreateAccountData.Unknown

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nintendoCreateAccountData *NintendoCreateAccountData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NintendoCreateAccountData)

	if !nintendoCreateAccountData.NNAInfo.Equals(other.NNAInfo) {
		return false
	}

	if nintendoCreateAccountData.Token != other.Token {
		return false
	}

	if !nintendoCreateAccountData.Birthday.Equals(other.Birthday) {
		return false
	}

	if nintendoCreateAccountData.Unknown != other.Unknown {
		return false
	}

	return true
}

// NewNintendoCreateAccountData returns a new NintendoCreateAccountData
func NewNintendoCreateAccountData() *NintendoCreateAccountData {
	return &NintendoCreateAccountData{}
}
