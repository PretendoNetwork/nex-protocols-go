// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

// GetCachedTopXRanking sets the GetCachedTopXRanking handler function
func (protocol *Protocol) GetCachedTopXRanking(handler func(err error, client *nex.Client, callID uint32, category uint32, orderParam *ranking_types.RankingOrderParam) uint32) {
	protocol.getCachedTopXRankingHandler = handler
}

func (protocol *Protocol) handleGetCachedTopXRanking(packet nex.PacketInterface) {
	if protocol.getCachedTopXRankingHandler == nil {
		globals.Logger.Warning("Ranking::GetCachedTopXRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getCachedTopXRankingHandler(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	orderParam, err := parametersStream.ReadStructure(ranking_types.NewRankingOrderParam())
	if err != nil {
		go protocol.getCachedTopXRankingHandler(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.getCachedTopXRankingHandler(nil, client, callID, category, orderParam.(*ranking_types.RankingOrderParam))
}
