// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetBufferQueue sets the GetBufferQueue handler function
func (protocol *Protocol) GetBufferQueue(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.BufferQueueParam) uint32) {
	protocol.getBufferQueueHandler = handler
}

func (protocol *Protocol) handleGetBufferQueue(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getBufferQueueHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetBufferQueue not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewBufferQueueParam())
	if err != nil {
		errorCode = protocol.getBufferQueueHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getBufferQueueHandler(nil, packet, callID, param.(*datastore_super_mario_maker_types.BufferQueueParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
