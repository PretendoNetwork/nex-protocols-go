// Package protocol implements the Uplay Win protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSetActionsCompleted(packet nex.PacketInterface) {
	if protocol.SetActionsCompleted == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UplayWinProtocol::SetActionsCompleted not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var actionCodes types.List[types.String]
	var platformCode types.String
	var gameCode types.String

	if err := actionCodes.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SetActionsCompleted(fmt.Errorf("failed to read actionCodes from parameters. %s", err.Error()), packet, callID, actionCodes, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	if err := platformCode.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SetActionsCompleted(fmt.Errorf("failed to read platformCode from parameters. %s", err.Error()), packet, callID, actionCodes, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	if err := gameCode.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SetActionsCompleted(fmt.Errorf("failed to read gameCode from parameters. %s", err.Error()), packet, callID, actionCodes, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SetActionsCompleted(nil, packet, callID, actionCodes, platformCode, gameCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
