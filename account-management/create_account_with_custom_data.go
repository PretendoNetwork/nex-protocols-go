// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CreateAccountWithCustomData sets the CreateAccountWithCustomData handler function
func (protocol *Protocol) CreateAccountWithCustomData(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) uint32) {
	protocol.createAccountWithCustomDataHandler = handler
}

func (protocol *Protocol) handleCreateAccountWithCustomData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.createAccountWithCustomDataHandler == nil {
		globals.Logger.Warning("AccountManagement::CreateAccountWithCustomData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strPrincipalName, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.createAccountWithCustomDataHandler(fmt.Errorf("Failed to read strPrincipalName from parameters. %s", err.Error()), packet, callID, "", "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strKey, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.createAccountWithCustomDataHandler(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), packet, callID, "", "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiGroups, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.createAccountWithCustomDataHandler(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), packet, callID, "", "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strEmail, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.createAccountWithCustomDataHandler(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), packet, callID, "", "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	oPublicData, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.createAccountWithCustomDataHandler(fmt.Errorf("Failed to read oPublicData from parameters. %s", err.Error()), packet, callID, "", "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	oPrivateData, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.createAccountWithCustomDataHandler(fmt.Errorf("Failed to read oPrivateData from parameters. %s", err.Error()), packet, callID, "", "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.createAccountWithCustomDataHandler(nil, packet, callID, strPrincipalName, strKey, uiGroups, strEmail, oPublicData, oPrivateData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
