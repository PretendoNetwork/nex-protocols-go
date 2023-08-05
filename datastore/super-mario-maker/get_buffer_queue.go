// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetBufferQueue sets the GetBufferQueue handler function
func (protocol *Protocol) GetBufferQueue(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.BufferQueueParam) uint32) {
	protocol.getBufferQueueHandler = handler
}

func (protocol *Protocol) handleGetBufferQueue(packet nex.PacketInterface) {
	if protocol.getBufferQueueHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetBufferQueue not implemented")
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
		go protocol.getBufferQueueHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getBufferQueueHandler(nil, client, callID, param.(*datastore_super_mario_maker_types.BufferQueueParam))
}
