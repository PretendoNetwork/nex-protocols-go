// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

// GetApproxOrder sets the GetApproxOrder handler function
func (protocol *Protocol) GetApproxOrder(handler func(err error, packet nex.PacketInterface, callID uint32, category uint32, orderParam *ranking_types.RankingOrderParam, score uint32, uniqueID uint64, principalID uint32) uint32) {
	protocol.getApproxOrderHandler = handler
}

func (protocol *Protocol) handleGetApproxOrder(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getApproxOrderHandler == nil {
		globals.Logger.Warning("Ranking::GetApproxOrder not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getApproxOrderHandler(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	orderParam, err := parametersStream.ReadStructure(ranking_types.NewRankingOrderParam())
	if err != nil {
		errorCode = protocol.getApproxOrderHandler(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	score, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getApproxOrderHandler(fmt.Errorf("Failed to read score from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.getApproxOrderHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	principalID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getApproxOrderHandler(fmt.Errorf("Failed to read principalID from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getApproxOrderHandler(nil, packet, callID, category, orderParam.(*ranking_types.RankingOrderParam), score, uniqueID, principalID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
