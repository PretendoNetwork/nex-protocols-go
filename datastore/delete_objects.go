// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteObjects(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.DeleteObjects == nil {
		globals.Logger.Warning("DataStore::DeleteObjects not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreDeleteParam())
	if err != nil {
		_, errorCode = protocol.DeleteObjects(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	transactional, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.DeleteObjects(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), packet, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DeleteObjects(nil, packet, callID, params.([]*datastore_types.DataStoreDeleteParam), transactional)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
