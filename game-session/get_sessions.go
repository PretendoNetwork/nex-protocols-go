// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/game-session/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetSessions(packet nex.PacketInterface) {
	if protocol.GetSessions == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "GameSession::GetSessions not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var gameSessionKeys types.List[game_session_types.GameSessionKey]

	err := gameSessionKeys.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetSessions(fmt.Errorf("failed to read gameSessionKeys from parameters. %s", err.Error()), packet, callID, gameSessionKeys)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetSessions(nil, packet, callID, gameSessionKeys)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
