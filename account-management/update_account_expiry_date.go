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
	var errorCode uint32

	if protocol.UpdateAccountExpiryDate == nil {
		globals.Logger.Warning("AccountManagement::UpdateAccountExpiryDate not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	idPrincipal := types.NewPID(0)
	err = idPrincipal.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateAccountExpiryDate(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dtExpiry := types.NewDateTime(0)
	err = dtExpiry.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateAccountExpiryDate(fmt.Errorf("Failed to read dtExpiry from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strExpiredMessage := types.NewString("")
	err = strExpiredMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateAccountExpiryDate(fmt.Errorf("Failed to read strExpiredMessage from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateAccountExpiryDate(nil, packet, callID, idPrincipal, dtExpiry, strExpiredMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
