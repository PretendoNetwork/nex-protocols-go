package account_management

import (
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

	accountExtraInfo.Unknown = stream.ReadUInt32LE()
	accountExtraInfo.Unknown2 = stream.ReadUInt32LE()
	accountExtraInfo.Unknown3 = stream.ReadUInt32LE()
	accountExtraInfo.NEXToken, _ = stream.ReadString()

	return nil
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
	nnaInfoStructureInterface, err := stream.ReadStructure(friends_wiiu.NewNNAInfo())
	if err != nil {
		return err
	}

	nnaInfo := nnaInfoStructureInterface.(*friends_wiiu.NNAInfo)

	token, err := stream.ReadString()
	if err != nil {
		return err
	}

	birthday := nex.NewDateTime(stream.ReadUInt64LE())
	unknown := stream.ReadUInt64LE()

	nintendoCreateAccountData.NNAInfo = nnaInfo
	nintendoCreateAccountData.Token = token
	nintendoCreateAccountData.Birthday = birthday
	nintendoCreateAccountData.Unknown = unknown

	return nil
}

// NewNintendoCreateAccountData returns a new NintendoCreateAccountData
func NewNintendoCreateAccountData() *NintendoCreateAccountData {
	return &NintendoCreateAccountData{}
}
