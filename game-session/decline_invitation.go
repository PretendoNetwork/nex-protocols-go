// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/game-session/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleDeclineInvitation(packet nex.PacketInterface) {
	if protocol.DeclineInvitation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "GameSession::DeclineInvitation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var gameSessionInvitation game_session_types.GameSessionInvitationReceived

	err := gameSessionInvitation.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeclineInvitation(fmt.Errorf("failed to read gameSessionInvitation from parameters. %s", err.Error()), packet, callID, gameSessionInvitation)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeclineInvitation(nil, packet, callID, gameSessionInvitation)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
