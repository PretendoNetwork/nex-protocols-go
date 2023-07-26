// Package matchmake_referee implements the Matchmake Referee NEX protocol
package matchmake_referee

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// EndRoundWithoutReport sets the EndRoundWithoutReport handler function
func (protocol *MatchmakeRefereeProtocol) EndRoundWithoutReport(handler func(err error, client *nex.Client, callID uint32, roundId uint64)) {
	protocol.EndRoundWithoutReportHandler = handler
}

func (protocol *MatchmakeRefereeProtocol) handleEndRoundWithoutReport(packet nex.PacketInterface) {
	if protocol.EndRoundWithoutReportHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::EndRoundWithoutReport not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	roundID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.EndRoundWithoutReportHandler(fmt.Errorf("Failed to read roundID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.EndRoundWithoutReportHandler(nil, client, callID, roundID)
}
