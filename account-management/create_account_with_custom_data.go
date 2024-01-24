// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleCreateAccountWithCustomData(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.CreateAccountWithCustomData == nil {
		globals.Logger.Warning("AccountManagement::CreateAccountWithCustomData not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	strPrincipalName := types.NewString("")
	err = strPrincipalName.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.CreateAccountWithCustomData(fmt.Errorf("Failed to read strPrincipalName from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strKey := types.NewString("")
	err = strKey.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.CreateAccountWithCustomData(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiGroups := types.NewPrimitiveU32(0)
	err = uiGroups.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.CreateAccountWithCustomData(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strEmail := types.NewString("")
	err = strEmail.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.CreateAccountWithCustomData(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	oPublicData := types.NewAnyDataHolder()
	err = oPublicData.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.CreateAccountWithCustomData(fmt.Errorf("Failed to read oPublicData from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	oPrivateData := types.NewAnyDataHolder()
	err = oPrivateData.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.CreateAccountWithCustomData(fmt.Errorf("Failed to read oPrivateData from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.CreateAccountWithCustomData(nil, packet, callID, strPrincipalName, strKey, uiGroups, strEmail, oPublicData, oPrivateData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
