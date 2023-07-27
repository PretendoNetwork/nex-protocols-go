// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ChangeMeta sets the ChangeMeta handler function
func (protocol *Protocol) ChangeMeta(handler func(err error, client *nex.Client, callID uint32, dataStoreChangeMetaParam *datastore_types.DataStoreChangeMetaParam)) {
	protocol.changeMetaHandler = handler
}

func (protocol *Protocol) handleChangeMeta(packet nex.PacketInterface) {
	if protocol.changeMetaHandler == nil {
		globals.Logger.Warning("DataStore::ChangeMeta not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreChangeMetaParam())
	if err != nil {
		go protocol.changeMetaHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.changeMetaHandler(nil, client, callID, param.(*datastore_types.DataStoreChangeMetaParam))
}
