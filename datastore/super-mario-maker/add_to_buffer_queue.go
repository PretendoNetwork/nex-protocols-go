// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddToBufferQueue(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AddToBufferQueue == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::AddToBufferQueue not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewBufferQueueParam())
	if err != nil {
		errorCode = protocol.AddToBufferQueue(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	buffer, err := parametersStream.ReadQBuffer()
	if err != nil {
		errorCode = protocol.AddToBufferQueue(fmt.Errorf("Failed to read buffer from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.AddToBufferQueue(nil, packet, callID, param.(*datastore_super_mario_maker_types.BufferQueueParam), buffer)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
