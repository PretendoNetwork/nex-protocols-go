package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareUpdateObject sets the PrepareUpdateObject handler function
func (protocol *DataStoreProtocol) PrepareUpdateObject(handler func(err error, client *nex.Client, callID uint32, dataStorePrepareUpdateParam *datastore_types.DataStorePrepareUpdateParam)) {
	protocol.PrepareUpdateObjectHandler = handler
}

func (protocol *DataStoreProtocol) handlePrepareUpdateObject(packet nex.PacketInterface) {
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

	dataStorePrepareUpdateParam, err := parametersStream.ReadStructure(datastore_types.NewDataStorePrepareUpdateParam())
	if err != nil {
		go protocol.PrepareUpdateObjectHandler(fmt.Errorf("Failed to read dataStorePrepareUpdateParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.PrepareUpdateObjectHandler(nil, client, callID, dataStorePrepareUpdateParam.(*datastore_types.DataStorePrepareUpdateParam))
}
