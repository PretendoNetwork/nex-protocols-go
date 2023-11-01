// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompleteAttachFileV1 sets the CompleteAttachFileV1 handler function
func (protocol *Protocol) CompleteAttachFileV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParamV1) uint32) {
	protocol.completeAttachFileV1Handler = handler
}

func (protocol *Protocol) handleCompleteAttachFileV1(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.completeAttachFileV1Handler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::CompleteAttachFileV1 not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompletePostParamV1())
	if err != nil {
		errorCode = protocol.completeAttachFileV1Handler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.completeAttachFileV1Handler(nil, packet, callID, param.(*datastore_types.DataStoreCompletePostParamV1))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
