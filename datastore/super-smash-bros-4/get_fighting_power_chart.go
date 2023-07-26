// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFightingPowerChart sets the GetFightingPowerChart handler function
func (protocol *Protocol) GetFightingPowerChart(handler func(err error, client *nex.Client, callID uint32, mode uint8)) {
	protocol.GetFightingPowerChartHandler = handler
}

func (protocol *Protocol) handleGetFightingPowerChart(packet nex.PacketInterface) {
	if protocol.GetFightingPowerChartHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::GetFightingPowerChart not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	mode, err := parametersStream.ReadUInt8()
	if err != nil {
		go protocol.GetFightingPowerChartHandler(fmt.Errorf("Failed to read mode from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.GetFightingPowerChartHandler(nil, client, callID, mode)
}
