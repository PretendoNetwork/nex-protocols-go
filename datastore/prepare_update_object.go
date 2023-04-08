package datastore

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareUpdateObject sets the PrepareUpdateObject handler function
func (protocol *DataStoreProtocol) PrepareUpdateObject(handler func(err error, client *nex.Client, callID uint32, dataStorePrepareUpdateParam *DataStorePrepareUpdateParam)) {
	protocol.PrepareUpdateObjectHandler = handler
}

func (protocol *DataStoreProtocol) HandlePrepareUpdateObject(packet nex.PacketInterface) {
	if protocol.PrepareUpdateObjectHandler == nil {
		globals.Logger.Warning("DataStore::PrepareUpdateObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStorePrepareUpdateParam, err := parametersStream.ReadStructure(NewDataStorePrepareUpdateParam())
	if err != nil {
		go protocol.PrepareUpdateObjectHandler(err, client, callID, nil)
		return
	}

	go protocol.PrepareUpdateObjectHandler(nil, client, callID, dataStorePrepareUpdateParam.(*DataStorePrepareUpdateParam))
}
