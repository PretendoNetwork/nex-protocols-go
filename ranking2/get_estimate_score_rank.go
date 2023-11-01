// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

// GetEstimateScoreRank sets the GetEstimateScoreRank handler function
func (protocol *Protocol) GetEstimateScoreRank(handler func(err error, packet nex.PacketInterface, callID uint32, input *ranking2_types.Ranking2EstimateScoreRankInput) uint32) {
	protocol.getEstimateScoreRankHandler = handler
}

func (protocol *Protocol) handleGetEstimateScoreRank(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getEstimateScoreRankHandler == nil {
		globals.Logger.Warning("Ranking2::GetEstimateScoreRank not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	input, err := parametersStream.ReadStructure(ranking2_types.NewRanking2EstimateScoreRankInput())
	if err != nil {
		errorCode = protocol.getEstimateScoreRankHandler(fmt.Errorf("Failed to read input from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getEstimateScoreRankHandler(nil, packet, callID, input.(*ranking2_types.Ranking2EstimateScoreRankInput))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
