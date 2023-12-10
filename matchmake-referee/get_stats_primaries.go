// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

func (protocol *Protocol) handleGetStatsPrimaries(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetStatsPrimaries == nil {
		globals.Logger.Warning("MatchmakeReferee::GetStatsPrimaries not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	targets, err := nex.StreamReadListStructure(parametersStream, matchmake_referee_types.NewMatchmakeRefereeStatsTarget())
	if err != nil {
		_, errorCode = protocol.GetStatsPrimaries(fmt.Errorf("Failed to read targets from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetStatsPrimaries(nil, packet, callID, targets)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
