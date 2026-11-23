// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/game-session/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSplitSession(packet nex.PacketInterface) {
	if protocol.SplitSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "GameSession::SplitSession not implemented")

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

	err := gameSessionKey.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SplitSession(fmt.Errorf("failed to read gameSessionKey from parameters. %s", err.Error()), packet, callID, gameSessionKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SplitSession(nil, packet, callID, gameSessionKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
