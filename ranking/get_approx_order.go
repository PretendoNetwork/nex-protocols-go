// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

func (protocol *Protocol) handleGetApproxOrder(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetApproxOrder == nil {
		globals.Logger.Warning("Ranking::GetApproxOrder not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	category := types.NewPrimitiveU32(0)
	err = category.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetApproxOrder(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	orderParam := ranking_types.NewRankingOrderParam()
	err = orderParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetApproxOrder(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	score := types.NewPrimitiveU32(0)
	err = score.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetApproxOrder(fmt.Errorf("Failed to read score from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uniqueID := types.NewPrimitiveU64(0)
	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetApproxOrder(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	principalID := types.NewPID(0)
	err = principalID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetApproxOrder(fmt.Errorf("Failed to read principalID from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetApproxOrder(nil, packet, callID, category, orderParam, score, uniqueID, principalID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
