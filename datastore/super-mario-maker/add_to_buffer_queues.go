// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddToBufferQueues sets the AddToBufferQueues handler function
func (protocol *Protocol) AddToBufferQueues(handler func(err error, packet nex.PacketInterface, callID uint32, params []*datastore_super_mario_maker_types.BufferQueueParam, buffers [][]byte) uint32) {
	protocol.addToBufferQueuesHandler = handler
}

func (protocol *Protocol) handleAddToBufferQueues(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.addToBufferQueuesHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::AddToBufferQueues not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_super_mario_maker_types.NewBufferQueueParam())
	if err != nil {
		errorCode = protocol.addToBufferQueuesHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	buffers, err := parametersStream.ReadListQBuffer()
	if err != nil {
		errorCode = protocol.addToBufferQueuesHandler(fmt.Errorf("Failed to read buffers from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.addToBufferQueuesHandler(nil, packet, callID, params.([]*datastore_super_mario_maker_types.BufferQueueParam), buffers)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
