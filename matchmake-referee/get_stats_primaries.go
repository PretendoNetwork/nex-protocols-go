// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

func (protocol *Protocol) handleGetStatsPrimaries(packet nex.PacketInterface) {
	var err error

	if protocol.GetStatsPrimaries == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeReferee::GetStatsPrimaries not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	targets := types.NewList[*matchmake_referee_types.MatchmakeRefereeStatsTarget]()
	targets.Type = matchmake_referee_types.NewMatchmakeRefereeStatsTarget()
	err = targets.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetStatsPrimaries(fmt.Errorf("Failed to read targets from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetStatsPrimaries(nil, packet, callID, targets)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
