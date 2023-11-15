// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

func (protocol *Protocol) handleGetEstimateScoreRank(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetEstimateScoreRank == nil {
		globals.Logger.Warning("Ranking2::GetEstimateScoreRank not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	input, err := parametersStream.ReadStructure(ranking2_types.NewRanking2EstimateScoreRankInput())
	if err != nil {
		_, errorCode = protocol.GetEstimateScoreRank(fmt.Errorf("Failed to read input from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetEstimateScoreRank(nil, packet, callID, input.(*ranking2_types.Ranking2EstimateScoreRankInput))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
