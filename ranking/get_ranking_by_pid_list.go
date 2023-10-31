// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

// GetRankingByPIDList sets the GetRankingByPIDList handler function
func (protocol *Protocol) GetRankingByPIDList(handler func(err error, packet nex.PacketInterface, callID uint32, principalIDList []uint32, rankingMode uint8, category uint32, orderParam *ranking_types.RankingOrderParam, uniqueID uint64) uint32) {
	protocol.getRankingByPIDListHandler = handler
}

func (protocol *Protocol) handleGetRankingByPIDList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getRankingByPIDListHandler == nil {
		globals.Logger.Warning("Ranking::GetRankingByPIDList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	principalIDList, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getRankingByPIDListHandler(fmt.Errorf("Failed to read principalIDList from parameters. %s", err.Error()), packet, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rankingMode, err := parametersStream.ReadUInt8()
	if err != nil {
		errorCode = protocol.getRankingByPIDListHandler(fmt.Errorf("Failed to read rankingMode from parameters. %s", err.Error()), packet, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getRankingByPIDListHandler(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	orderParam, err := parametersStream.ReadStructure(ranking_types.NewRankingOrderParam())
	if err != nil {
		errorCode = protocol.getRankingByPIDListHandler(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.getRankingByPIDListHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getRankingByPIDListHandler(nil, packet, callID, principalIDList, rankingMode, category, orderParam.(*ranking_types.RankingOrderParam), uniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
