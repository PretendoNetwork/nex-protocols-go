// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSendPlayReport(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.SendPlayReport == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::SendPlayReport not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	playReport := types.NewList[*types.PrimitiveS32]()
	playReport.Type = types.NewPrimitiveS32(0)
	err = playReport.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SendPlayReport(fmt.Errorf("Failed to read playReport from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.SendPlayReport(nil, packet, callID, playReport)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
