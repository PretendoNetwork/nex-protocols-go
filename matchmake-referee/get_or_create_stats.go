// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

func (protocol *Protocol) handleGetOrCreateStats(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetOrCreateStats == nil {
		globals.Logger.Warning("MatchmakeReferee::GetOrCreateStats not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(matchmake_referee_types.NewMatchmakeRefereeStatsInitParam())
	if err != nil {
		errorCode = protocol.GetOrCreateStats(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.GetOrCreateStats(nil, packet, callID, param.(*matchmake_referee_types.MatchmakeRefereeStatsInitParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
