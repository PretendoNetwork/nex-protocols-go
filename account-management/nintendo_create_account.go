package account_management

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// NintendoCreateAccount sets the NintendoCreateAccount handler function
func (protocol *AccountManagementProtocol) NintendoCreateAccount(handler func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder)) {
	protocol.NintendoCreateAccountHandler = handler
}

func (protocol *AccountManagementProtocol) HandleNintendoCreateAccount(packet nex.PacketInterface) {
	if protocol.NintendoCreateAccountHandler == nil {
		globals.Logger.Warning("AccountManagement::NintendoCreateAccount not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strPrincipalName, err := parametersStream.ReadString()
	if err != nil {
		go protocol.NintendoCreateAccountHandler(err, client, callID, "", "", 0, "", nil)
		return
	}

	strKey, err := parametersStream.ReadString()
	if err != nil {
		go protocol.NintendoCreateAccountHandler(err, client, callID, "", "", 0, "", nil)
		return
	}

	uiGroups := parametersStream.ReadUInt32LE()
	strEmail, err := parametersStream.ReadString()
	if err != nil {
		go protocol.NintendoCreateAccountHandler(err, client, callID, "", "", 0, "", nil)
		return
	}

	oAuthData := parametersStream.ReadDataHolder()

	go protocol.NintendoCreateAccountHandler(nil, client, callID, strPrincipalName, strKey, uiGroups, strEmail, oAuthData)
}
