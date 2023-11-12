// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

// GetRankingChart sets the GetRankingChart handler function
func (protocol *Protocol) GetRankingChart(handler func(err error, packet nex.PacketInterface, callID uint32, info *ranking2_types.Ranking2ChartInfoInput) uint32) {
	protocol.getRankingChartHandler = handler
}

func (protocol *Protocol) handleGetRankingChart(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getRankingChartHandler == nil {
		globals.Logger.Warning("Ranking2::GetRankingChart not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	info, err := parametersStream.ReadStructure(ranking2_types.NewRanking2ChartInfoInput())
	if err != nil {
		errorCode = protocol.getRankingChartHandler(fmt.Errorf("Failed to read info from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getRankingChartHandler(nil, packet, callID, info.(*ranking2_types.Ranking2ChartInfoInput))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
