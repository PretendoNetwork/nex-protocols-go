// Package datastore implements the DataStore NEX protocol
package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetas sets the GetMetas handler function
func (protocol *DataStoreProtocol) GetMetas(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64, param *datastore_types.DataStoreGetMetaParam)) {
	protocol.GetMetasHandler = handler
}

func (protocol *DataStoreProtocol) handleGetMetas(packet nex.PacketInterface) {
	if protocol.GetMetasHandler == nil {
		globals.Logger.Warning("DataStore::GetMetas not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDs, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.GetMetasHandler(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetMetaParam())
	if err != nil {
		go protocol.GetMetasHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.GetMetasHandler(nil, client, callID, dataIDs, param.(*datastore_types.DataStoreGetMetaParam))
}
