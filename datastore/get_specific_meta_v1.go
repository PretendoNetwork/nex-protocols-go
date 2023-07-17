// Package datastore implements the DataStore NEX protocol
package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSpecificMetaV1 sets the GetSpecificMetaV1 handler function
func (protocol *DataStoreProtocol) GetSpecificMetaV1(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParamV1)) {
	protocol.GetSpecificMetaV1Handler = handler
}

func (protocol *DataStoreProtocol) handleGetSpecificMetaV1(packet nex.PacketInterface) {
	if protocol.PrepareGetObjectV1Handler == nil {
		globals.Logger.Warning("DataStore::GetSpecificMetaV1 not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetSpecificMetaParamV1())
	if err != nil {
		go protocol.GetSpecificMetaV1Handler(fmt.Errorf("Failed to read dataStoreGetSpecificMetaParamV1 from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetSpecificMetaV1Handler(nil, client, callID, param.(*datastore_types.DataStoreGetSpecificMetaParamV1))
}
