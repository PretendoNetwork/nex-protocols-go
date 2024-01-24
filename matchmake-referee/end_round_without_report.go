// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleEndRoundWithoutReport(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.EndRoundWithoutReport == nil {
		globals.Logger.Warning("MatchmakeReferee::EndRoundWithoutReport not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	roundID := types.NewPrimitiveU64(0)
	err = roundID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.EndRoundWithoutReport(fmt.Errorf("Failed to read roundID from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.EndRoundWithoutReport(nil, packet, callID, roundID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
