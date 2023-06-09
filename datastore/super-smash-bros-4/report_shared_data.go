package datastore_super_smash_bros_4

import (
	"fmt"

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

	dataID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.ReportSharedDataHandler(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.ReportSharedDataHandler(nil, client, callID, dataID)
}
