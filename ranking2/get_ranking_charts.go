// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

// GetRankingCharts sets the GetRankingCharts handler function
func (protocol *Protocol) GetRankingCharts(handler func(err error, packet nex.PacketInterface, callID uint32, infoArray []*ranking2_types.Ranking2ChartInfoInput) uint32) {
	protocol.getRankingChartsHandler = handler
}

func (protocol *Protocol) handleGetRankingCharts(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getRankingChartsHandler == nil {
		globals.Logger.Warning("Ranking2::GetRankingCharts not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	infoArray, err := parametersStream.ReadListStructure(ranking2_types.NewRanking2ChartInfoInput())
	if err != nil {
		errorCode = protocol.getRankingChartsHandler(fmt.Errorf("Failed to read infoArray from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getRankingChartsHandler(nil, packet, callID, infoArray.([]*ranking2_types.Ranking2ChartInfoInput))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
