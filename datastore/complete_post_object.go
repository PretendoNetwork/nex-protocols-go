// Package datastore implements the DataStore NEX protocol
package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostObject sets the CompletePostObject handler function
func (protocol *DataStoreProtocol) CompletePostObject(handler func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParam *datastore_types.DataStoreCompletePostParam)) {
	protocol.CompletePostObjectHandler = handler
}

func (protocol *DataStoreProtocol) handleCompletePostObject(packet nex.PacketInterface) {
	if protocol.CompletePostObjectHandler == nil {
		globals.Logger.Warning("DataStore::CompletePostObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStoreCompletePostParam, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompletePostParam())
	if err != nil {
		go protocol.CompletePostObjectHandler(fmt.Errorf("Failed to read dataStoreCompletePostParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.CompletePostObjectHandler(nil, client, callID, dataStoreCompletePostParam.(*datastore_types.DataStoreCompletePostParam))
}
