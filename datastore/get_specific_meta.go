// Package datastore implements the DataStore NEX protocol
package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSpecificMeta sets the GetSpecificMeta handler function
func (protocol *DataStoreProtocol) GetSpecificMeta(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParam)) {
	protocol.GetSpecificMetaHandler = handler
}

func (protocol *DataStoreProtocol) handleGetSpecificMeta(packet nex.PacketInterface) {
	if protocol.GetSpecificMetaHandler == nil {
		globals.Logger.Warning("DataStore::GetSpecificMeta not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetSpecificMetaParam())
	if err != nil {
		go protocol.GetSpecificMetaHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetSpecificMetaHandler(nil, client, callID, param.(*datastore_types.DataStoreGetSpecificMetaParam))
}
