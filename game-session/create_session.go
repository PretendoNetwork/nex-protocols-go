// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/game-session/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleCreateSession(packet nex.PacketInterface) {
	if protocol.CreateSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "GameSession::CreateSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var gameSession game_session_types.GameSession

	err := gameSession.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateSession(fmt.Errorf("failed to read gameSession from parameters. %s", err.Error()), packet, callID, gameSession)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.CreateSession(nil, packet, callID, gameSession)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
