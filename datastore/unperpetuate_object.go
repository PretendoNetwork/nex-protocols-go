// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUnperpetuateObject(packet nex.PacketInterface) {
	if protocol.UnperpetuateObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::UnperpetuateObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	persistenceSlotID := types.NewPrimitiveU16(0)
	deleteLastObject := types.NewPrimitiveBool(false)

	var err error

	err = persistenceSlotID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UnperpetuateObject(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = deleteLastObject.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UnperpetuateObject(fmt.Errorf("Failed to read deleteLastObject from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UnperpetuateObject(nil, packet, callID, persistenceSlotID, deleteLastObject)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
