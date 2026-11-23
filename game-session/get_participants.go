// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/game-session/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetParticipants(packet nex.PacketInterface) {
	if protocol.GetParticipants == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "GameSession::GetParticipants not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var gameSessionKey game_session_types.GameSessionKey
	var resultRange types.ResultRange

	var err error

	err = gameSessionKey.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetParticipants(fmt.Errorf("failed to read gameSessionKey from parameters. %s", err.Error()), packet, callID, gameSessionKey, resultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetParticipants(fmt.Errorf("failed to read resultRange from parameters. %s", err.Error()), packet, callID, gameSessionKey, resultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetParticipants(nil, packet, callID, gameSessionKey, resultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
