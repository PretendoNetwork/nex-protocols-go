// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddToBufferQueue(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.AddToBufferQueue == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::AddToBufferQueue not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	param := datastore_super_mario_maker_types.NewBufferQueueParam()
	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AddToBufferQueue(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	buffer := types.NewQBuffer(nil)
	err = buffer.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AddToBufferQueue(fmt.Errorf("Failed to read buffer from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AddToBufferQueue(nil, packet, callID, param, buffer)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
