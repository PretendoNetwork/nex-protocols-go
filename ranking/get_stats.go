// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

func (protocol *Protocol) handleGetStats(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetStats == nil {
		globals.Logger.Warning("Ranking::GetStats not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetStats(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	orderParam, err := nex.StreamReadStructure(parametersStream, ranking_types.NewRankingOrderParam())
	if err != nil {
		_, errorCode = protocol.GetStats(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	flags, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetStats(fmt.Errorf("Failed to read flags from parameters. %s", err.Error()), packet, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetStats(nil, packet, callID, category, orderParam, flags)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
