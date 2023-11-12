// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// EndRoundWithoutReport sets the EndRoundWithoutReport handler function
func (protocol *Protocol) EndRoundWithoutReport(handler func(err error, packet nex.PacketInterface, callID uint32, roundId uint64) uint32) {
	protocol.endRoundWithoutReportHandler = handler
}

func (protocol *Protocol) handleEndRoundWithoutReport(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.endRoundWithoutReportHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::EndRoundWithoutReport not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	roundID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.endRoundWithoutReportHandler(fmt.Errorf("Failed to read roundID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.endRoundWithoutReportHandler(nil, packet, callID, roundID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
