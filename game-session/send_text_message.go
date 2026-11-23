// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/game-session/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSendTextMessage(packet nex.PacketInterface) {
	if protocol.SendTextMessage == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "GameSession::SendTextMessage not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var gameSessionMessage game_session_types.GameSessionMessage

	err := gameSessionMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SendTextMessage(fmt.Errorf("failed to read gameSessionMessage from parameters. %s", err.Error()), packet, callID, gameSessionMessage)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SendTextMessage(nil, packet, callID, gameSessionMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
