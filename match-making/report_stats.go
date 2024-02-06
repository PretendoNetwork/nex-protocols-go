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
	if protocol.ReportStats == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::ReportStats not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	idGathering := types.NewPrimitiveU32(0)
	lstStats := types.NewList[*match_making_types.GatheringStats]()
	lstStats.Type = match_making_types.NewGatheringStats()

	var err error

	err = idGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportStats(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = lstStats.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportStats(fmt.Errorf("Failed to read lstStats from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReportStats(nil, packet, callID, idGathering, lstStats)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
