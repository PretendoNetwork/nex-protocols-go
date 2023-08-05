// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddToBufferQueue sets the AddToBufferQueue handler function
func (protocol *Protocol) AddToBufferQueue(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.BufferQueueParam, buffer []byte) uint32) {
	protocol.addToBufferQueueHandler = handler
}

func (protocol *Protocol) handleAddToBufferQueue(packet nex.PacketInterface) {
	if protocol.addToBufferQueueHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::AddToBufferQueue not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewBufferQueueParam())
	if err != nil {
		go protocol.addToBufferQueueHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	buffer, err := parametersStream.ReadQBuffer()
	if err != nil {
		go protocol.addToBufferQueueHandler(fmt.Errorf("Failed to read buffer from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.addToBufferQueueHandler(nil, client, callID, param.(*datastore_super_mario_maker_types.BufferQueueParam), buffer)
}
