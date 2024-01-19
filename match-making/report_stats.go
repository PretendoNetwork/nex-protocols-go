// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func (protocol *Protocol) handleReportStats(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.ReportStats == nil {
		globals.Logger.Warning("MatchMaking::ReportStats not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	idGathering := types.NewPrimitiveU32(0)
	err = idGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ReportStats(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstStats := types.NewList[*match_making_types.GatheringStats]()
	lstStats.Type = match_making_types.NewGatheringStats()
	err = lstStats.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ReportStats(fmt.Errorf("Failed to read lstStats from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.ReportStats(nil, packet, callID, idGathering, lstStats)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
