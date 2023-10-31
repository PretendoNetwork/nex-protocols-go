// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFightingPowerChartAll sets the GetFightingPowerChartAll handler function
func (protocol *Protocol) GetFightingPowerChartAll(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.getFightingPowerChartAllHandler = handler
}

func (protocol *Protocol) handleGetFightingPowerChartAll(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getFightingPowerChartAllHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetFightingPowerChartAll not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getFightingPowerChartAllHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
