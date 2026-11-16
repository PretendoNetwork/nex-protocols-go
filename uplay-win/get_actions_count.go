// Package protocol implements the Uplay Win protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetActionsCount(packet nex.PacketInterface) {
	if protocol.GetActionsCount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UplayWinProtocol::GetActionsCount not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var platformCode types.String
	var gameCode types.String

	if err := platformCode.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetActionsCount(fmt.Errorf("failed to read platformCode from parameters. %s", err.Error()), packet, callID, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}
		return
	}

	if err := gameCode.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetActionsCount(fmt.Errorf("failed to read gameCode from parameters. %s", err.Error()), packet, callID, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}
		return
	}

	rmcMessage, rmcError := protocol.GetActionsCount(nil, packet, callID, platformCode, gameCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
