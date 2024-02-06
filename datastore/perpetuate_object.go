// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handlePerpetuateObject(packet nex.PacketInterface) {
	var err error

	if protocol.PerpetuateObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::PerpetuateObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	persistenceSlotID := types.NewPrimitiveU16(0)
	err = persistenceSlotID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PerpetuateObject(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	dataID := types.NewPrimitiveU64(0)
	err = dataID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PerpetuateObject(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	deleteLastObject := types.NewPrimitiveBool(false)
	err = deleteLastObject.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PerpetuateObject(fmt.Errorf("Failed to read deleteLastObject from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PerpetuateObject(nil, packet, callID, persistenceSlotID, dataID, deleteLastObject)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
