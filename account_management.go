package nexproto

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// AccountManagementProtocolID is the protocol ID for the Account Management protocol
	AccountManagementProtocolID = 0x19

	// AccountManagementMethodNintendoCreateAccount is the method ID for the method NintendoCreateAccount
	AccountManagementMethodNintendoCreateAccount = 0x1B
)

// AccountManagementProtocol handles the Account Management nex protocol
type AccountManagementProtocol struct {
	server                          *nex.Server
	NintendoCreateAccountHandler    func(err error, client *nex.Client, callID uint32, username string, key string, groups uint32, email string, nintendoCreateAccountData *NintendoCreateAccountData)
	NintendoCreateAccount3DSHandler func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *AccountExtraInfo)
}

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
	NNAInfo  *NNAInfo
	Token    string
	Birthday *nex.DateTime
	Unknown  uint64
}

// ExtractFromStream extracts a NintendoCreateAccountData structure from a stream
func (nintendoCreateAccountData *NintendoCreateAccountData) ExtractFromStream(stream *nex.StreamIn) error {
	nnaInfoStructureInterface, err := stream.ReadStructure(NewNNAInfo())
	if err != nil {
		return err
	}

	nnaInfo := nnaInfoStructureInterface.(*NNAInfo)

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

// Setup initializes the protocol
func (accountManagementProtocol *AccountManagementProtocol) Setup() {
	nexServer := accountManagementProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if AccountManagementProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case AccountManagementMethodNintendoCreateAccount:
				go accountManagementProtocol.handleNintendoCreateAccountHandler(packet)
			default:
				go respondNotImplemented(packet, AccountManagementProtocolID)
				fmt.Printf("Unsupported AccountManagement method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NintendoCreateAccount sets the NintendoCreateAccount handler function
func (accountManagementProtocol *AccountManagementProtocol) NintendoCreateAccount(handler func(err error, client *nex.Client, callID uint32, username string, key string, groups uint32, email string, nintendoCreateAccountData *NintendoCreateAccountData)) {
	accountManagementProtocol.NintendoCreateAccountHandler = handler
}

// NintendoCreateAccount sets the NintendoCreateAccount handler function. This is not a real NEX method, this was made specifically for nex-protocols-go to simplify things
func (accountManagementProtocol *AccountManagementProtocol) NintendoCreateAccount3DS(handler func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *AccountExtraInfo)) {
	accountManagementProtocol.NintendoCreateAccount3DSHandler = handler
}

func (accountManagementProtocol *AccountManagementProtocol) handleNintendoCreateAccountHandler(packet nex.PacketInterface) {
	if accountManagementProtocol.NintendoCreateAccountHandler == nil {
		fmt.Println("[Warning] AccountManagementProtocol::NintendoCreateAccount not implemented")
		go respondNotImplemented(packet, AccountManagementProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, accountManagementProtocol.server)

	username, err := parametersStream.ReadString()
	if err != nil {
		go accountManagementProtocol.NintendoCreateAccountHandler(err, client, callID, "", "", 0, "", nil)
		return
	}

	key, err := parametersStream.ReadString()
	if err != nil {
		go accountManagementProtocol.NintendoCreateAccountHandler(err, client, callID, "", "", 0, "", nil)
		return
	}

	groups := parametersStream.ReadUInt32LE()
	email, err := parametersStream.ReadString()
	if err != nil {
		go accountManagementProtocol.NintendoCreateAccountHandler(err, client, callID, "", "", 0, "", nil)
		return
	}

	dataHolderName, err := parametersStream.ReadString()
	if err != nil {
		go accountManagementProtocol.NintendoCreateAccountHandler(err, client, callID, "", "", 0, "", nil)
		return
	}

	if dataHolderName != "NintendoCreateAccountData" && dataHolderName != "AccountExtraInfo" {
		err := errors.New("[AccountManagementProtocol::NintendoCreateAccount] Data holder name does not match")
		go accountManagementProtocol.NintendoCreateAccountHandler(err, client, callID, "", "", 0, "", nil)
		return
	}

	_ = parametersStream.ReadUInt32LE() // length including this field

	dataHolderContent, err := parametersStream.ReadBuffer()
	if err != nil {
		go accountManagementProtocol.NintendoCreateAccountHandler(err, client, callID, "", "", 0, "", nil)
		return
	}

	dataHolderContentStream := nex.NewStreamIn(dataHolderContent, accountManagementProtocol.server)

	if dataHolderName == "NintendoCreateAccountData" {
		nintendoCreateAccountDataStructureInterface, err := dataHolderContentStream.ReadStructure(NewNintendoCreateAccountData())
		if err != nil {
			go accountManagementProtocol.NintendoCreateAccountHandler(err, client, callID, "", "", 0, "", nil)
			return
		}

		nintendoCreateAccountData := nintendoCreateAccountDataStructureInterface.(*NintendoCreateAccountData)

		go accountManagementProtocol.NintendoCreateAccountHandler(nil, client, callID, username, key, groups, email, nintendoCreateAccountData)
	}

	if dataHolderName == "AccountExtraInfo" {
		accountExtraInfo, err := dataHolderContentStream.ReadStructure(NewAccountExtraInfo())
		if err != nil {
			go accountManagementProtocol.NintendoCreateAccount3DSHandler(err, client, callID, "", "", 0, "", nil)
			return
		}

		go accountManagementProtocol.NintendoCreateAccount3DSHandler(nil, client, callID, username, key, groups, email, accountExtraInfo.(*AccountExtraInfo))
	}
}

// NewAccountManagementProtocol returns a new AccountManagementProtocol
func NewAccountManagementProtocol(server *nex.Server) *AccountManagementProtocol {
	accountManagementProtocol := &AccountManagementProtocol{server: server}

	accountManagementProtocol.Setup()

	return accountManagementProtocol
}
