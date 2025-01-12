// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleAddToBufferQueues(packet nex.PacketInterface) {
	if protocol.AddToBufferQueues == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::AddToBufferQueues not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params types.List[datastore_super_mario_maker_types.BufferQueueParam]
	var buffers types.List[types.QBuffer]

	var err error

	err = params.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddToBufferQueues(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, params, buffers)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = buffers.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddToBufferQueues(fmt.Errorf("Failed to read buffers from parameters. %s", err.Error()), packet, callID, params, buffers)
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
