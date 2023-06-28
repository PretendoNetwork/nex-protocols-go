package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMeta sets the GetMeta handler function
func (protocol *DataStoreProtocol) GetMeta(handler func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParam *datastore_types.DataStoreGetMetaParam)) {
	protocol.GetMetaHandler = handler
}

func (protocol *DataStoreProtocol) HandleGetMeta(packet nex.PacketInterface) {
	if protocol.GetMetaHandler == nil {
		globals.Logger.Warning("DataStore::GetMeta not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStoreGetMetaParam, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetMetaParam())
	if err != nil {
		go protocol.GetMetaHandler(fmt.Errorf("Failed to read dataStoreGetMetaParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetMetaHandler(nil, client, callID, dataStoreGetMetaParam.(*datastore_types.DataStoreGetMetaParam))
}
