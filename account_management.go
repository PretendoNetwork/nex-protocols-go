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
	server                       *nex.Server
	NintendoCreateAccountHandler func(err error, client *nex.Client, callID uint32, username string, key string, groups uint32, email string, nintendoCreateAccountData *NintendoCreateAccountData)
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
				fmt.Printf("Unsupported AccountManagement method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NintendoCreateAccount sets the NintendoCreateAccount handler function
func (accountManagementProtocol *AccountManagementProtocol) NintendoCreateAccount(handler func(err error, client *nex.Client, callID uint32, username string, key string, groups uint32, email string, nintendoCreateAccountData *NintendoCreateAccountData)) {
	accountManagementProtocol.NintendoCreateAccountHandler = handler
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

	if dataHolderName != "NintendoCreateAccountData" {
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

	nintendoCreateAccountDataStructureInterface, err := dataHolderContentStream.ReadStructure(NewNintendoCreateAccountData())
	if err != nil {
		go accountManagementProtocol.NintendoCreateAccountHandler(err, client, callID, "", "", 0, "", nil)
		return
	}

	nintendoCreateAccountData := nintendoCreateAccountDataStructureInterface.(*NintendoCreateAccountData)

	go accountManagementProtocol.NintendoCreateAccountHandler(nil, client, callID, username, key, groups, email, nintendoCreateAccountData)
}

// NewAccountManagementProtocol returns a new AccountManagementProtocol
func NewAccountManagementProtocol(server *nex.Server) *AccountManagementProtocol {
	accountManagementProtocol := &AccountManagementProtocol{server: server}

	accountManagementProtocol.Setup()

	return accountManagementProtocol
}
