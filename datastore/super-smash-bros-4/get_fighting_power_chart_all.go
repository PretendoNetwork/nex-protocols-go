// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFightingPowerChartAll sets the GetFightingPowerChartAll handler function
func (protocol *Protocol) GetFightingPowerChartAll(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.getFightingPowerChartAllHandler = handler
}

func (protocol *Protocol) handleGetFightingPowerChartAll(packet nex.PacketInterface) {
	if protocol.getFightingPowerChartAllHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetFightingPowerChartAll not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getFightingPowerChartAllHandler(nil, client, callID)
}
