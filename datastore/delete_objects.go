// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteObjects sets the DeleteObjects handler function
func (protocol *Protocol) DeleteObjects(handler func(err error, client *nex.Client, callID uint32, params []*datastore_types.DataStoreDeleteParam, transactional bool) uint32) {
	protocol.deleteObjectsHandler = handler
}

func (protocol *Protocol) handleDeleteObjects(packet nex.PacketInterface) {
	if protocol.deleteObjectsHandler == nil {
		globals.Logger.Warning("DataStore::DeleteObjects not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreDeleteParam())
	if err != nil {
		go protocol.deleteObjectsHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil, false)
		return
	}

	transactional, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.deleteObjectsHandler(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), client, callID, nil, false)
		return
	}

	go protocol.deleteObjectsHandler(nil, client, callID, params.([]*datastore_types.DataStoreDeleteParam), transactional)
}
