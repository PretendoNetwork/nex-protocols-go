// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// ReportStats sets the ReportStats handler function
func (protocol *Protocol) ReportStats(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, lstStats []*match_making_types.GatheringStats) uint32) {
	protocol.reportStatsHandler = handler
}

func (protocol *Protocol) handleReportStats(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.reportStatsHandler == nil {
		globals.Logger.Warning("MatchMaking::ReportStats not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.reportStatsHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstStats, err := parametersStream.ReadListStructure(match_making_types.NewGatheringStats())
	if err != nil {
		errorCode = protocol.reportStatsHandler(fmt.Errorf("Failed to read lstStats from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.reportStatsHandler(nil, packet, callID, idGathering, lstStats.([]*match_making_types.GatheringStats))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
