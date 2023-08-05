// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CreateAccount sets the CreateAccount handler function
func (protocol *Protocol) CreateAccount(handler func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string) uint32) {
	protocol.createAccountHandler = handler
}

func (protocol *Protocol) handleCreateAccount(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.createAccountHandler == nil {
		globals.Logger.Warning("AccountManagement::CreateAccount not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strPrincipalName, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.createAccountHandler(fmt.Errorf("Failed to read strPrincipalName from parameters. %s", err.Error()), client, callID, "", "", 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strKey, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.createAccountHandler(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), client, callID, "", "", 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiGroups, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.createAccountHandler(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), client, callID, "", "", 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strEmail, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.createAccountHandler(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), client, callID, "", "", 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.createAccountHandler(nil, client, callID, strPrincipalName, strKey, uiGroups, strEmail)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
