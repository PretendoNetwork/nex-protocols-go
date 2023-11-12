// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostObject sets the CompletePostObject handler function
func (protocol *Protocol) CompletePostObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParam) uint32) {
	protocol.completePostObjectHandler = handler
}

func (protocol *Protocol) handleCompletePostObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.completePostObjectHandler == nil {
		globals.Logger.Warning("DataStore::CompletePostObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompletePostParam())
	if err != nil {
		errorCode = protocol.completePostObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.completePostObjectHandler(nil, packet, callID, param.(*datastore_types.DataStoreCompletePostParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
