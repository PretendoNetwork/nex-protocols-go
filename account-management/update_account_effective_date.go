// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateAccountEffectiveDate(packet nex.PacketInterface) {
	var err error

	if protocol.UpdateAccountEffectiveDate == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::UpdateAccountEffectiveDate not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	idPrincipal := types.NewPID(0)
	err = idPrincipal.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAccountEffectiveDate(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	dtEffectiveFrom := types.NewDateTime(0)
	err = dtEffectiveFrom.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAccountEffectiveDate(fmt.Errorf("Failed to read dtEffectiveFrom from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	strNotEffectiveMessage := types.NewString("")
	err = strNotEffectiveMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAccountEffectiveDate(fmt.Errorf("Failed to read strNotEffectiveMessage from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateAccountEffectiveDate(nil, packet, callID, idPrincipal, dtEffectiveFrom, strNotEffectiveMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
