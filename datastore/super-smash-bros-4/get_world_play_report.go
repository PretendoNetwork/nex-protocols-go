package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetWorldPlayReport sets the GetWorldPlayReport handler function
func (protocol *DataStoreSuperSmashBros4Protocol) GetWorldPlayReport(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetWorldPlayReportHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) handleGetWorldPlayReport(packet nex.PacketInterface) {
	if protocol.GetWorldPlayReportHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::GetWorldPlayReport not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.GetWorldPlayReportHandler(nil, client, callID)
}
