// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

// GetCachedTopXRankings sets the GetCachedTopXRankings handler function
func (protocol *Protocol) GetCachedTopXRankings(handler func(err error, client *nex.Client, callID uint32, categories []uint32, orderParams []*ranking_types.RankingOrderParam) uint32) {
	protocol.getCachedTopXRankingsHandler = handler
}

func (protocol *Protocol) handleGetCachedTopXRankings(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getCachedTopXRankingsHandler == nil {
		globals.Logger.Warning("Ranking::GetCachedTopXRankings not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	categories, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getCachedTopXRankingsHandler(fmt.Errorf("Failed to read categories from parameters. %s", err.Error()), client, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	orderParams, err := parametersStream.ReadListStructure(ranking_types.NewRankingOrderParam())
	if err != nil {
		errorCode = protocol.getCachedTopXRankingsHandler(fmt.Errorf("Failed to read orderParams from parameters. %s", err.Error()), client, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getCachedTopXRankingsHandler(nil, client, callID, categories, orderParams.([]*ranking_types.RankingOrderParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
