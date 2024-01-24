// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetFightingPowerChart(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetFightingPowerChart == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetFightingPowerChart not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	mode := types.NewPrimitiveU8(0)
	err = mode.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetFightingPowerChart(fmt.Errorf("Failed to read mode from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetFightingPowerChart(nil, packet, callID, mode)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
