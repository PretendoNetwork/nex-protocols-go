// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

// GetRankingByUniqueIDList sets the GetRankingByUniqueIDList handler function
func (protocol *Protocol) GetRankingByUniqueIDList(handler func(err error, client *nex.Client, callID uint32, nexUniqueIDList []uint64, rankingMode uint8, category uint32, orderParam *ranking_types.RankingOrderParam, uniqueID uint64) uint32) {
	protocol.getRankingByUniqueIDListHandler = handler
}

func (protocol *Protocol) handleGetRankingByUniqueIDList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getRankingByUniqueIDListHandler == nil {
		globals.Logger.Warning("Ranking::GetRankingByUniqueIDList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	nexUniqueIDList, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.getRankingByUniqueIDListHandler(fmt.Errorf("Failed to read nexUniqueIDList from parameters. %s", err.Error()), client, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rankingMode, err := parametersStream.ReadUInt8()
	if err != nil {
		errorCode = protocol.getRankingByUniqueIDListHandler(fmt.Errorf("Failed to read rankingMode from parameters. %s", err.Error()), client, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getRankingByUniqueIDListHandler(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), client, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	orderParam, err := parametersStream.ReadStructure(ranking_types.NewRankingOrderParam())
	if err != nil {
		errorCode = protocol.getRankingByUniqueIDListHandler(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), client, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.getRankingByUniqueIDListHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getRankingByUniqueIDListHandler(nil, client, callID, nexUniqueIDList, rankingMode, category, orderParam.(*ranking_types.RankingOrderParam), uniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
