// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetEstimateScoreRank sets the GetEstimateScoreRank handler function
func (protocol *Protocol) GetEstimateScoreRank(handler func(err error, client *nex.Client, callID uint32, input *ranking2_types.Ranking2EstimateScoreRankInput)) {
	protocol.getEstimateScoreRankHandler = handler
}

func (protocol *Protocol) handleGetEstimateScoreRank(packet nex.PacketInterface) {
	if protocol.getEstimateScoreRankHandler == nil {
		globals.Logger.Warning("Ranking2::GetEstimateScoreRank not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	input, err := parametersStream.ReadStructure(ranking2_types.NewRanking2EstimateScoreRankInput())
	if err != nil {
		go protocol.getEstimateScoreRankHandler(fmt.Errorf("Failed to read input from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getEstimateScoreRankHandler(nil, client, callID, input.(*ranking2_types.Ranking2EstimateScoreRankInput))
}