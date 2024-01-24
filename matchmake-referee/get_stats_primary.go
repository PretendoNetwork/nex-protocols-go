// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

func (protocol *Protocol) handleGetStatsPrimary(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetStatsPrimary == nil {
		globals.Logger.Warning("MatchmakeReferee::GetStatsPrimary not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	target := matchmake_referee_types.NewMatchmakeRefereeStatsTarget()
	err = target.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetStatsPrimary(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetStatsPrimary(nil, packet, callID, target)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
