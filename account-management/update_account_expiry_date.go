// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateAccountExpiryDate(packet nex.PacketInterface) {
	var err error

	if protocol.UpdateAccountExpiryDate == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::UpdateAccountExpiryDate not implemented")

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
		_, rmcError := protocol.UpdateAccountExpiryDate(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	dtExpiry := types.NewDateTime(0)
	err = dtExpiry.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAccountExpiryDate(fmt.Errorf("Failed to read dtExpiry from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	strExpiredMessage := types.NewString("")
	err = strExpiredMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAccountExpiryDate(fmt.Errorf("Failed to read strExpiredMessage from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateAccountExpiryDate(nil, packet, callID, idPrincipal, dtExpiry, strExpiredMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
