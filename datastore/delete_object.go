// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteObject sets the DeleteObject handler function
func (protocol *Protocol) DeleteObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreDeleteParam)) {
	protocol.deleteObjectHandler = handler
}

func (protocol *Protocol) handleDeleteObject(packet nex.PacketInterface) {
	if protocol.deleteObjectHandler == nil {
		globals.Logger.Warning("DataStore::DeleteObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreDeleteParam())
	if err != nil {
		go protocol.deleteObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.deleteObjectHandler(nil, client, callID, param.(*datastore_types.DataStoreDeleteParam))
}
