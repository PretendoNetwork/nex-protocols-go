package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFightingPowerChartAll sets the GetFightingPowerChartAll handler function
func (protocol *DataStoreSuperSmashBros4Protocol) GetFightingPowerChartAll(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetFightingPowerChartAllHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandleGetFightingPowerChartAll(packet nex.PacketInterface) {
	if protocol.GetFightingPowerChartAllHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::GetFightingPowerChartAll not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.GetFightingPowerChartAllHandler(nil, client, callID)
}
