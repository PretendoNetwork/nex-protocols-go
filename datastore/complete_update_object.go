// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompleteUpdateObject sets the CompleteUpdateObject handler function
func (protocol *Protocol) CompleteUpdateObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompleteUpdateParam) uint32) {
	protocol.completeUpdateObjectHandler = handler
}

func (protocol *Protocol) handleCompleteUpdateObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.completeUpdateObjectHandler == nil {
		globals.Logger.Warning("DataStore::CompleteUpdateObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompleteUpdateParam())
	if err != nil {
		errorCode = protocol.completeUpdateObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.completeUpdateObjectHandler(nil, packet, callID, param.(*datastore_types.DataStoreCompleteUpdateParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
