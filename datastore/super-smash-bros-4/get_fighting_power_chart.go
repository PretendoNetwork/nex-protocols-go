package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFightingPowerChart sets the GetFightingPowerChart handler function
func (protocol *DataStoreSuperSmashBros4Protocol) GetFightingPowerChart(handler func(err error, client *nex.Client, callID uint32, mode uint8)) {
	protocol.GetFightingPowerChartHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandleGetFightingPowerChart(packet nex.PacketInterface) {
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

	mode := parametersStream.ReadUInt8()

	go protocol.GetFightingPowerChartHandler(nil, client, callID, mode)
}