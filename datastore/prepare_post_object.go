package datastore

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PreparePostObject sets the PreparePostObject handler function
func (protocol *DataStoreProtocol) PreparePostObject(handler func(err error, client *nex.Client, callID uint32, dataStorePreparePostParam *DataStorePreparePostParam)) {
	protocol.PreparePostObjectHandler = handler
}

func (protocol *DataStoreProtocol) HandlePreparePostObject(packet nex.PacketInterface) {
	if protocol.PreparePostObjectHandler == nil {
		globals.Logger.Warning("DataStore::PreparePostObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStorePreparePostParam, err := parametersStream.ReadStructure(NewDataStorePreparePostParam())
	if err != nil {
		go protocol.PreparePostObjectHandler(err, client, callID, nil)
		return
	}

	go protocol.PreparePostObjectHandler(nil, client, callID, dataStorePreparePostParam.(*DataStorePreparePostParam))
}
