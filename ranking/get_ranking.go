// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

// GetRanking sets the GetRanking handler function
func (protocol *Protocol) GetRanking(handler func(err error, client *nex.Client, callID uint32, rankingMode uint8, category uint32, orderParam *ranking_types.RankingOrderParam, uniqueID uint64, principalID uint32)) {
	protocol.GetRankingHandler = handler
}

func (protocol *Protocol) handleGetRanking(packet nex.PacketInterface) {
	if protocol.GetRankingHandler == nil {
		globals.Logger.Warning("Ranking::GetRanking not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	rankingMode, err := parametersStream.ReadUInt8()
	if err != nil {
		go protocol.GetRankingHandler(fmt.Errorf("Failed to read rankingMode from parameters. %s", err.Error()), client, callID, 0, 0, nil, 0, 0)
		return
	}

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.GetRankingHandler(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), client, callID, 0, 0, nil, 0, 0)
		return
	}

	orderParam, err := parametersStream.ReadStructure(ranking_types.NewRankingOrderParam())
	if err != nil {
		go protocol.GetRankingHandler(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), client, callID, 0, 0, nil, 0, 0)
		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.GetRankingHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, 0, 0, nil, 0, 0)
		return
	}

	principalID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.GetRankingHandler(fmt.Errorf("Failed to read principalID from parameters. %s", err.Error()), client, callID, 0, 0, nil, 0, 0)
		return
	}

	go protocol.GetRankingHandler(nil, client, callID, rankingMode, category, orderParam.(*ranking_types.RankingOrderParam), uniqueID, principalID)
}
