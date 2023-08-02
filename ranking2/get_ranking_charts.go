// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRankingCharts sets the GetRankingCharts handler function
func (protocol *Protocol) GetRankingCharts(handler func(err error, client *nex.Client, callID uint32, infoArray []*ranking2_types.Ranking2ChartInfoInput)) {
	protocol.getRankingChartsHandler = handler
}

func (protocol *Protocol) handleGetRankingCharts(packet nex.PacketInterface) {
	if protocol.getRankingChartsHandler == nil {
		globals.Logger.Warning("Ranking2::GetRankingCharts not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	infoArray, err := parametersStream.ReadListStructure(ranking2_types.NewRanking2ChartInfoInput())
	if err != nil {
		go protocol.getRankingChartsHandler(fmt.Errorf("Failed to read infoArray from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getRankingChartsHandler(nil, client, callID, infoArray.([]*ranking2_types.Ranking2ChartInfoInput))
}