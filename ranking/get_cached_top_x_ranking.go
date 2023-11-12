// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

// GetCachedTopXRanking sets the GetCachedTopXRanking handler function
func (protocol *Protocol) GetCachedTopXRanking(handler func(err error, packet nex.PacketInterface, callID uint32, category uint32, orderParam *ranking_types.RankingOrderParam) uint32) {
	protocol.getCachedTopXRankingHandler = handler
}

func (protocol *Protocol) handleGetCachedTopXRanking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getCachedTopXRankingHandler == nil {
		globals.Logger.Warning("Ranking::GetCachedTopXRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getCachedTopXRankingHandler(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	orderParam, err := parametersStream.ReadStructure(ranking_types.NewRankingOrderParam())
	if err != nil {
		errorCode = protocol.getCachedTopXRankingHandler(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getCachedTopXRankingHandler(nil, packet, callID, category, orderParam.(*ranking_types.RankingOrderParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
