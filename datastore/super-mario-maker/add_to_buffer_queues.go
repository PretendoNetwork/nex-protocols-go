// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddToBufferQueues(packet nex.PacketInterface) {
	var err error

	if protocol.AddToBufferQueues == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::AddToBufferQueues not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	params := types.NewList[*datastore_super_mario_maker_types.BufferQueueParam]()
	params.Type = datastore_super_mario_maker_types.NewBufferQueueParam()
	err = params.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddToBufferQueues(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	buffers := types.NewList[*types.QBuffer]()
	buffers.Type = types.NewQBuffer(nil)
	err = buffers.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddToBufferQueues(fmt.Errorf("Failed to read buffers from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AddToBufferQueues(nil, packet, callID, params, buffers)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
