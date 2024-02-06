// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDisconnectPrincipal(packet nex.PacketInterface) {
	if protocol.DisconnectPrincipal == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::DisconnectPrincipal not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	idPrincipal := types.NewPID(0)

	err := idPrincipal.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DisconnectPrincipal(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DisconnectPrincipal(nil, packet, callID, idPrincipal)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
