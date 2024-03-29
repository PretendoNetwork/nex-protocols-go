// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSpecificMeta sets the GetSpecificMeta handler function
func (protocol *Protocol) GetSpecificMeta(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParam) uint32) {
	protocol.getSpecificMetaHandler = handler
}

func (protocol *Protocol) handleGetSpecificMeta(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getSpecificMetaHandler == nil {
		globals.Logger.Warning("DataStore::GetSpecificMeta not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetSpecificMetaParam())
	if err != nil {
		errorCode = protocol.getSpecificMetaHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getSpecificMetaHandler(nil, packet, callID, param.(*datastore_types.DataStoreGetSpecificMetaParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
