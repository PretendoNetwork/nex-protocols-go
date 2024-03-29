// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFightingPowerChart sets the GetFightingPowerChart handler function
func (protocol *Protocol) GetFightingPowerChart(handler func(err error, packet nex.PacketInterface, callID uint32, mode uint8) uint32) {
	protocol.getFightingPowerChartHandler = handler
}

func (protocol *Protocol) handleGetFightingPowerChart(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getFightingPowerChartHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetFightingPowerChart not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	mode, err := parametersStream.ReadUInt8()
	if err != nil {
		errorCode = protocol.getFightingPowerChartHandler(fmt.Errorf("Failed to read mode from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getFightingPowerChartHandler(nil, packet, callID, mode)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
