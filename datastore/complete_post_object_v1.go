// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostObjectV1 sets the CompletePostObjectV1 handler function
func (protocol *Protocol) CompletePostObjectV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParamV1) uint32) {
	protocol.completePostObjectV1Handler = handler
}

func (protocol *Protocol) handleCompletePostObjectV1(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.completePostObjectV1Handler == nil {
		globals.Logger.Warning("DataStore::CompletePostObjectV1 not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompletePostParamV1())
	if err != nil {
		errorCode = protocol.completePostObjectV1Handler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.completePostObjectV1Handler(nil, packet, callID, param.(*datastore_types.DataStoreCompletePostParamV1))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
