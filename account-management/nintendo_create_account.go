// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleNintendoCreateAccount(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.NintendoCreateAccount == nil {
		globals.Logger.Warning("AccountManagement::NintendoCreateAccount not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	strPrincipalName := types.NewString("")
	err = strPrincipalName.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.NintendoCreateAccount(fmt.Errorf("Failed to read strPrincipalName from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strKey := types.NewString("")
	err = strKey.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.NintendoCreateAccount(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiGroups := types.NewPrimitiveU32(0)
	err = uiGroups.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.NintendoCreateAccount(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strEmail := types.NewString("")
	err = strEmail.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.NintendoCreateAccount(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	oAuthData := types.NewAnyDataHolder()
	err = oAuthData.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.NintendoCreateAccount(fmt.Errorf("Failed to read oAuthData from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.NintendoCreateAccount(nil, packet, callID, strPrincipalName, strKey, uiGroups, strEmail, oAuthData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
