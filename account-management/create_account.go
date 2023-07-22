// Package account_management implements the Account Management NEX protocol
package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CreateAccount sets the CreateAccount handler function
func (protocol *AccountManagementProtocol) CreateAccount(handler func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string)) {
	protocol.createAccountHandler = handler
}

func (protocol *AccountManagementProtocol) handleCreateAccount(packet nex.PacketInterface) {
	if protocol.createAccountHandler == nil {
		globals.Logger.Warning("AccountManagement::CreateAccount not implemented")
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
		go protocol.createAccountHandler(fmt.Errorf("Failed to read strPrincipalName from parameters. %s", err.Error()), client, callID, "", "", 0, "")
		return
	}

	strKey, err := parametersStream.ReadString()
	if err != nil {
		go protocol.createAccountHandler(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), client, callID, "", "", 0, "")
		return
	}

	uiGroups, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.createAccountHandler(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), client, callID, "", "", 0, "")
		return
	}

	strEmail, err := parametersStream.ReadString()
	if err != nil {
		go protocol.createAccountHandler(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), client, callID, "", "", 0, "")
		return
	}

	go protocol.createAccountHandler(nil, client, callID, strPrincipalName, strKey, uiGroups, strEmail)
}
