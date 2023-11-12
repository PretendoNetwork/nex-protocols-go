// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMeta sets the GetMeta handler function
func (protocol *Protocol) GetMeta(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetMetaParam) uint32) {
	protocol.getMetaHandler = handler
}

func (protocol *Protocol) handleGetMeta(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getMetaHandler == nil {
		globals.Logger.Warning("DataStore::GetMeta not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetMetaParam())
	if err != nil {
		errorCode = protocol.getMetaHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getMetaHandler(nil, packet, callID, param.(*datastore_types.DataStoreGetMetaParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
