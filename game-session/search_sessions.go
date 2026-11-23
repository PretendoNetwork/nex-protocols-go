// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/game-session/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSearchSessions(packet nex.PacketInterface) {
	if protocol.SearchSessions == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "GameSession::SearchSessions not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var gameSessionQuery game_session_types.GameSessionQuery

	err := gameSessionQuery.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SearchSessions(fmt.Errorf("failed to read gameSessionQuery from parameters. %s", err.Error()), packet, callID, gameSessionQuery)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchSessions(nil, packet, callID, gameSessionQuery)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
