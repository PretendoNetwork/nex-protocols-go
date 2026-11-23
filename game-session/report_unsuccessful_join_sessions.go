// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/game-session/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleReportUnsuccessfulJoinSessions(packet nex.PacketInterface) {
	if protocol.ReportUnsuccessfulJoinSessions == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "GameSession::ReportUnsuccessfulJoinSessions not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var unsuccessfulJoinSessions types.List[game_session_types.GameSessionUnsuccessfulJoinSession]

	err := unsuccessfulJoinSessions.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportUnsuccessfulJoinSessions(fmt.Errorf("failed to read unsuccessfulJoinSessions from parameters. %s", err.Error()), packet, callID, unsuccessfulJoinSessions)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReportUnsuccessfulJoinSessions(nil, packet, callID, unsuccessfulJoinSessions)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
