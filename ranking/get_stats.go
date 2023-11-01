// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

// GetStats sets the GetStats handler function
func (protocol *Protocol) GetStats(handler func(err error, packet nex.PacketInterface, callID uint32, category uint32, orderParam *ranking_types.RankingOrderParam, flags uint32) uint32) {
	protocol.getStatsHandler = handler
}

func (protocol *Protocol) handleGetStats(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getStatsHandler == nil {
		globals.Logger.Warning("Ranking::GetStats not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getStatsHandler(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	orderParam, err := parametersStream.ReadStructure(ranking_types.NewRankingOrderParam())
	if err != nil {
		errorCode = protocol.getStatsHandler(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	flags, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getStatsHandler(fmt.Errorf("Failed to read flags from parameters. %s", err.Error()), packet, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getStatsHandler(nil, packet, callID, category, orderParam.(*ranking_types.RankingOrderParam), flags)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
