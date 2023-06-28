package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddToBufferQueues sets the AddToBufferQueues handler function
func (protocol *DataStoreSuperMarioMakerProtocol) AddToBufferQueues(handler func(err error, client *nex.Client, callID uint32, params []*datastore_super_mario_maker_types.BufferQueueParam, buffers [][]byte)) {
	protocol.AddToBufferQueuesHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleAddToBufferQueues(packet nex.PacketInterface) {
	if protocol.AddToBufferQueuesHandler == nil {
		globals.Logger.Warning("DataStoreSMM::AddToBufferQueues not implemented")
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
		go protocol.AddToBufferQueuesHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	buffers, err := parametersStream.ReadListQBuffer()
	if err != nil {
		go protocol.AddToBufferQueuesHandler(fmt.Errorf("Failed to read buffers from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.AddToBufferQueuesHandler(nil, client, callID, params.([]*datastore_super_mario_maker_types.BufferQueueParam), buffers)
}
