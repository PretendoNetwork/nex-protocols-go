// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

func (protocol *Protocol) handleGetCachedTopXRankings(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetCachedTopXRankings == nil {
		globals.Logger.Warning("Ranking::GetCachedTopXRankings not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	categories := types.NewList[*types.PrimitiveU32]()
	categories.Type = types.NewPrimitiveU32(0)
	err = categories.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetCachedTopXRankings(fmt.Errorf("Failed to read categories from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	orderParams := types.NewList[*ranking_types.RankingOrderParam]()
	orderParams.Type = ranking_types.NewRankingOrderParam()
	err = orderParams.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetCachedTopXRankings(fmt.Errorf("Failed to read orderParams from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetCachedTopXRankings(nil, packet, callID, categories, orderParams)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
