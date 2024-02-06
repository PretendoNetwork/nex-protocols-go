// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

func (protocol *Protocol) handleGetRankingByPIDList(packet nex.PacketInterface) {
	var err error

	if protocol.GetRankingByPIDList == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::GetRankingByPIDList not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	principalIDList := types.NewList[*types.PID]()
	principalIDList.Type = types.NewPID(0)
	err = principalIDList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPIDList(fmt.Errorf("Failed to read principalIDList from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rankingMode := types.NewPrimitiveU8(0)
	err = rankingMode.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPIDList(fmt.Errorf("Failed to read rankingMode from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	category := types.NewPrimitiveU32(0)
	err = category.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPIDList(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	orderParam := ranking_types.NewRankingOrderParam()
	err = orderParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPIDList(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	uniqueID := types.NewPrimitiveU64(0)
	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPIDList(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetRankingByPIDList(nil, packet, callID, principalIDList, rankingMode, category, orderParam, uniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
