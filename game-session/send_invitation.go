// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/game-session/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSendInvitation(packet nex.PacketInterface) {
	if protocol.SendInvitation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "GameSession::SendInvitation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var invitation game_session_types.GameSessionInvitation

	err := invitation.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SendInvitation(fmt.Errorf("failed to read invitation from parameters. %s", err.Error()), packet, callID, invitation)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SendInvitation(nil, packet, callID, invitation)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
