package datastore_super_mario_maker

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddToBufferQueues sets the AddToBufferQueues handler function
func (protocol *DataStoreSuperMarioMakerProtocol) AddToBufferQueues(handler func(err error, client *nex.Client, callID uint32, params []*BufferQueueParam, buffers [][]byte)) {
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

	params, err := parametersStream.ReadListStructure(NewBufferQueueParam())
	if err != nil {
		go protocol.AddToBufferQueuesHandler(err, client, callID, nil, nil)
		return
	}

	buffers := parametersStream.ReadListQBuffer()

	go protocol.AddToBufferQueuesHandler(nil, client, callID, params.([]*BufferQueueParam), buffers)
}
