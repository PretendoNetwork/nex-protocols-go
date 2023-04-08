package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportSharedData sets the ReportSharedData handler function
func (protocol *DataStoreSuperSmashBros4Protocol) ReportSharedData(handler func(err error, client *nex.Client, callID uint32, dataID uint64)) {
	protocol.ReportSharedDataHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandleReportSharedData(packet nex.PacketInterface) {
	if protocol.ReportSharedDataHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::ReportSharedData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataID := parametersStream.ReadUInt64LE()

	go protocol.ReportSharedDataHandler(nil, client, callID, dataID)
}
