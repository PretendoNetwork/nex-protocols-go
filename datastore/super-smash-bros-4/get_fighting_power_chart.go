// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetFightingPowerChart(packet nex.PacketInterface) {
	if protocol.GetFightingPowerChart == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperSmashBros4::GetFightingPowerChart not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	mode := types.NewPrimitiveU8(0)

	err := mode.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetFightingPowerChart(fmt.Errorf("Failed to read mode from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetFightingPowerChart(nil, packet, callID, mode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
