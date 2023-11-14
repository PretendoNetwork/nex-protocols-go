// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetFightingPowerChartAll(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetFightingPowerChartAll == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetFightingPowerChartAll not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.GetFightingPowerChartAll(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
