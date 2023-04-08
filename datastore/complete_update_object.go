package datastore

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompleteUpdateObject sets the CompleteUpdateObject handler function
func (protocol *DataStoreProtocol) CompleteUpdateObject(handler func(err error, client *nex.Client, callID uint32, dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam)) {
	protocol.CompleteUpdateObjectHandler = handler
}

func (protocol *DataStoreProtocol) HandleCompleteUpdateObject(packet nex.PacketInterface) {
	if protocol.CompleteUpdateObjectHandler == nil {
		globals.Logger.Warning("DataStore::CompleteUpdateObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStoreCompleteUpdateParam, err := parametersStream.ReadStructure(NewDataStoreCompleteUpdateParam())
	if err != nil {
		go protocol.CompleteUpdateObjectHandler(err, client, callID, nil)
		return
	}

	go protocol.CompleteUpdateObjectHandler(nil, client, callID, dataStoreCompleteUpdateParam.(*DataStoreCompleteUpdateParam))
}
