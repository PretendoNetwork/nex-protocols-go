// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

func (protocol *Protocol) handleGetRankingByPIDList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetRankingByPIDList == nil {
		globals.Logger.Warning("Ranking::GetRankingByPIDList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	principalIDList, err := parametersStream.ReadListPID()
	if err != nil {
		_, errorCode = protocol.GetRankingByPIDList(fmt.Errorf("Failed to read principalIDList from parameters. %s", err.Error()), packet, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rankingMode, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.GetRankingByPIDList(fmt.Errorf("Failed to read rankingMode from parameters. %s", err.Error()), packet, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetRankingByPIDList(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	orderParam, err := parametersStream.ReadStructure(ranking_types.NewRankingOrderParam())
	if err != nil {
		_, errorCode = protocol.GetRankingByPIDList(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		_, errorCode = protocol.GetRankingByPIDList(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, nil, 0, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetRankingByPIDList(nil, packet, callID, principalIDList, rankingMode, category, orderParam.(*ranking_types.RankingOrderParam), uniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
