package nexproto

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	AccountManagementProtocolID = 0x19

	AccountManagementMethodNintendoCreateAccount = 0x1B
)

type AccountManagementProtocol struct {
	server                       *nex.Server
	NintendoCreateAccountHandler func(err error, client *nex.Client, callID uint32, username string, key string, groups uint32, email string, nintendoCreateAccountData *NintendoCreateAccountData)
}

type NintendoCreateAccountData struct {
	nnaInfo  *NNAInfo
	token    string
	birthday *nex.DateTime
	unknown  uint64

	hierarchy []nex.StructureInterface
	nex.Structure
}

func (nintendoCreateAccountData *NintendoCreateAccountData) GetNNAInfo() *NNAInfo {
	return nintendoCreateAccountData.nnaInfo
}

func (nintendoCreateAccountData *NintendoCreateAccountData) GetToken() string {
	return nintendoCreateAccountData.token
}

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

	nintendoCreateAccountData.nnaInfo = nnaInfo
	nintendoCreateAccountData.token = token
	nintendoCreateAccountData.birthday = birthday
	nintendoCreateAccountData.unknown = unknown

	return nil
}

func NewNintendoCreateAccountData() *NintendoCreateAccountData {
	return &NintendoCreateAccountData{}
}

func (accountManagementProtocol *AccountManagementProtocol) Setup() {
	nexServer := accountManagementProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.GetRMCRequest()

		if AccountManagementProtocolID == request.GetProtocolID() {
			switch request.GetMethodID() {
			case AccountManagementMethodNintendoCreateAccount:
				go accountManagementProtocol.handleNintendoCreateAccountHandler(packet)
			default:
				fmt.Printf("Unsupported AccountManagement method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

func (accountManagementProtocol *AccountManagementProtocol) respondNotImplemented(packet nex.PacketInterface) {
	client := packet.GetSender()
	request := packet.GetRMCRequest()

	rmcResponse := nex.NewRMCResponse(AccountManagementProtocolID, request.GetCallID())
	rmcResponse.SetError(0x80010002)

	rmcResponseBytes := rmcResponse.Bytes()

	var responsePacket nex.PacketInterface
	if packet.GetVersion() == 1 {
		responsePacket, _ = nex.NewPacketV1(client, nil)
	} else {
		responsePacket, _ = nex.NewPacketV0(client, nil)
	}

	responsePacket.SetVersion(packet.GetVersion())
	responsePacket.SetSource(packet.GetDestination())
	responsePacket.SetDestination(packet.GetSource())
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	accountManagementProtocol.server.Send(responsePacket)
}

func (accountManagementProtocol *AccountManagementProtocol) NintendoCreateAccount(handler func(err error, client *nex.Client, callID uint32, username string, key string, groups uint32, email string, nintendoCreateAccountData *NintendoCreateAccountData)) {
	accountManagementProtocol.NintendoCreateAccountHandler = handler
}

func (accountManagementProtocol *AccountManagementProtocol) handleNintendoCreateAccountHandler(packet nex.PacketInterface) {
	if accountManagementProtocol.NintendoCreateAccountHandler == nil {
		fmt.Println("[Warning] AccountManagementProtocol::NintendoCreateAccount not implemented")
		go accountManagementProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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

func NewAccountManagementProtocol(server *nex.Server) *AccountManagementProtocol {
	accountManagementProtocol := &AccountManagementProtocol{server: server}

	accountManagementProtocol.Setup()

	return accountManagementProtocol
}
