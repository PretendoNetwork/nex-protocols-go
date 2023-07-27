// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetasMultipleParam sets the GetMetasMultipleParam handler function
func (protocol *Protocol) GetMetasMultipleParam(handler func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParams []*datastore_types.DataStoreGetMetaParam)) {
	protocol.getMetasMultipleParamHandler = handler
}

func (protocol *Protocol) handleGetMetasMultipleParam(packet nex.PacketInterface) {
	if protocol.getMetasMultipleParamHandler == nil {
		globals.Logger.Warning("DataStore::GetMetasMultipleParam not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreGetMetaParam())
	if err != nil {
		go protocol.getMetasMultipleParamHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getMetasMultipleParamHandler(nil, client, callID, params.([]*datastore_types.DataStoreGetMetaParam))
}
