// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

func (protocol *Protocol) handleGetApproxOrder(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetApproxOrder == nil {
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
		_, errorCode = protocol.GetApproxOrder(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	orderParam, err := nex.StreamReadStructure(parametersStream, ranking_types.NewRankingOrderParam())
	if err != nil {
		_, errorCode = protocol.GetApproxOrder(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	score, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetApproxOrder(fmt.Errorf("Failed to read score from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		_, errorCode = protocol.GetApproxOrder(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, 0, nil, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	principalID, err := parametersStream.ReadPID()
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
