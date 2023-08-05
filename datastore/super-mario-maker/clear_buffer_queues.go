// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ClearBufferQueues sets the ClearBufferQueues handler function
func (protocol *Protocol) ClearBufferQueues(handler func(err error, client *nex.Client, callID uint32, params []*datastore_super_mario_maker_types.BufferQueueParam) uint32) {
	protocol.clearBufferQueuesHandler = handler
}

func (protocol *Protocol) handleClearBufferQueues(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.clearBufferQueuesHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::ClearBufferQueues not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_super_mario_maker_types.NewBufferQueueParam())
	if err != nil {
		errorCode = protocol.clearBufferQueuesHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.clearBufferQueuesHandler(nil, client, callID, params.([]*datastore_super_mario_maker_types.BufferQueueParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
