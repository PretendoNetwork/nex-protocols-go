// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

func (protocol *Protocol) handleEndRound(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.EndRound == nil {
		globals.Logger.Warning("MatchmakeReferee::EndRound not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	endRoundParam, err := parametersStream.ReadStructure(matchmake_referee_types.NewMatchmakeRefereeEndRoundParam())
	if err != nil {
		errorCode = protocol.EndRound(fmt.Errorf("Failed to read endRoundParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.EndRound(nil, packet, callID, endRoundParam.(*matchmake_referee_types.MatchmakeRefereeEndRoundParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
