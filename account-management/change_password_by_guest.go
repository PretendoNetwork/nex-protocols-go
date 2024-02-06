// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleChangePasswordByGuest(packet nex.PacketInterface) {
	var err error

	if protocol.ChangePasswordByGuest == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::ChangePasswordByGuest not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	strPrincipalName := types.NewString("")
	err = strPrincipalName.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangePasswordByGuest(fmt.Errorf("Failed to read strPrincipalName from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	strKey := types.NewString("")
	err = strKey.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangePasswordByGuest(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	strEmail := types.NewString("")
	err = strEmail.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangePasswordByGuest(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ChangePasswordByGuest(nil, packet, callID, strPrincipalName, strKey, strEmail)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
