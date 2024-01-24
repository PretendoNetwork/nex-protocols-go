// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

func (protocol *Protocol) handleEndRound(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.EndRound == nil {
		globals.Logger.Warning("MatchmakeReferee::EndRound not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	endRoundParam := matchmake_referee_types.NewMatchmakeRefereeEndRoundParam()
	err = endRoundParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.EndRound(fmt.Errorf("Failed to read endRoundParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.EndRound(nil, packet, callID, endRoundParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
