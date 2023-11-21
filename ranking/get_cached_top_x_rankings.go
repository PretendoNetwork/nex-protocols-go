// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

func (protocol *Protocol) handleGetCachedTopXRankings(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetCachedTopXRankings == nil {
		globals.Logger.Warning("Ranking::GetCachedTopXRankings not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	categories, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetCachedTopXRankings(fmt.Errorf("Failed to read categories from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	orderParams, err := nex.StreamReadListStructure(parametersStream, ranking_types.NewRankingOrderParam())
	if err != nil {
		_, errorCode = protocol.GetCachedTopXRankings(fmt.Errorf("Failed to read orderParams from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetCachedTopXRankings(nil, packet, callID, categories, orderParams)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
