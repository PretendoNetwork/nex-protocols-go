package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompleteUpdateObject sets the CompleteUpdateObject handler function
func (protocol *DataStoreProtocol) CompleteUpdateObject(handler func(err error, client *nex.Client, callID uint32, dataStoreCompleteUpdateParam *datastore_types.DataStoreCompleteUpdateParam)) {
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

	dataStoreCompleteUpdateParam, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompleteUpdateParam())
	if err != nil {
		go protocol.CompleteUpdateObjectHandler(fmt.Errorf("Failed to read dataStoreCompleteUpdateParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.CompleteUpdateObjectHandler(nil, client, callID, dataStoreCompleteUpdateParam.(*datastore_types.DataStoreCompleteUpdateParam))
}
