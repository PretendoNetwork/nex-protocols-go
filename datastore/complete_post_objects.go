// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleCompletePostObjects(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.CompletePostObjects == nil {
		globals.Logger.Warning("DataStore::CompletePostObjects not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	dataIDs := types.NewList[*types.PrimitiveU64]()
	dataIDs.Type = types.NewPrimitiveU64(0)
	err = dataIDs.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.CompletePostObjects(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.CompletePostObjects(nil, packet, callID, dataIDs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
