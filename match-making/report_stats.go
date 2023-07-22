// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// ReportStats sets the ReportStats handler function
func (protocol *MatchMakingProtocol) ReportStats(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, lstStats []*match_making_types.GatheringStats)) {
	protocol.reportStatsHandler = handler
}

func (protocol *MatchMakingProtocol) handleReportStats(packet nex.PacketInterface) {
	if protocol.reportStatsHandler == nil {
		globals.Logger.Warning("MatchMaking::ReportStats not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.reportStatsHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, nil)
	}

	lstStats, err := parametersStream.ReadListStructure(match_making_types.NewGatheringStats())
	if err != nil {
		go protocol.reportStatsHandler(fmt.Errorf("Failed to read lstStats from parameters. %s", err.Error()), client, callID, 0, nil)
	}

	go protocol.reportStatsHandler(nil, client, callID, idGathering, lstStats.([]*match_making_types.GatheringStats))
}
