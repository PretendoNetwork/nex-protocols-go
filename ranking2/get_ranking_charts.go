// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

func (protocol *Protocol) handleGetRankingCharts(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetRankingCharts == nil {
		globals.Logger.Warning("Ranking2::GetRankingCharts not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	infoArray := types.NewList[*ranking2_types.Ranking2ChartInfoInput]()
	infoArray.Type = ranking2_types.NewRanking2ChartInfoInput()
	err = infoArray.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetRankingCharts(fmt.Errorf("Failed to read infoArray from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetRankingCharts(nil, packet, callID, infoArray)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
