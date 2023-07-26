// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// EndRoundWithoutReport sets the EndRoundWithoutReport handler function
func (protocol *Protocol) EndRoundWithoutReport(handler func(err error, client *nex.Client, callID uint32, roundId uint64)) {
	protocol.endRoundWithoutReportHandler = handler
}

func (protocol *Protocol) handleEndRoundWithoutReport(packet nex.PacketInterface) {
	if protocol.endRoundWithoutReportHandler == nil {
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
		go protocol.endRoundWithoutReportHandler(fmt.Errorf("Failed to read roundID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.endRoundWithoutReportHandler(nil, client, callID, roundID)
}
