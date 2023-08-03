// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddToBufferQueues sets the AddToBufferQueues handler function
func (protocol *Protocol) AddToBufferQueues(handler func(err error, client *nex.Client, callID uint32, params []*datastore_super_mario_maker_types.BufferQueueParam, buffers [][]byte)) {
	protocol.addToBufferQueuesHandler = handler
}

func (protocol *Protocol) handleAddToBufferQueues(packet nex.PacketInterface) {
	if protocol.addToBufferQueuesHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::AddToBufferQueues not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_super_mario_maker_types.NewBufferQueueParam())
	if err != nil {
		go protocol.addToBufferQueuesHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	buffers, err := parametersStream.ReadListQBuffer()
	if err != nil {
		go protocol.addToBufferQueuesHandler(fmt.Errorf("Failed to read buffers from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.addToBufferQueuesHandler(nil, client, callID, params.([]*datastore_super_mario_maker_types.BufferQueueParam), buffers)
}
