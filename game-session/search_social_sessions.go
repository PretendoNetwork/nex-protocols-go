// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/game-session/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSearchSocialSessions(packet nex.PacketInterface) {
	if protocol.SearchSocialSessions == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "GameSession::SearchSocialSessions not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var gameSessionSocialQuery game_session_types.GameSessionSocialQuery

	err := gameSessionSocialQuery.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SearchSocialSessions(fmt.Errorf("failed to read gameSessionSocialQuery from parameters. %s", err.Error()), packet, callID, gameSessionSocialQuery)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchSocialSessions(nil, packet, callID, gameSessionSocialQuery)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
