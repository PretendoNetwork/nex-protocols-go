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
	var errorCode uint32

	if protocol.PerpetuateObject == nil {
		globals.Logger.Warning("DataStore::PerpetuateObject not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	persistenceSlotID := types.NewPrimitiveU16(0)
	err = persistenceSlotID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.PerpetuateObject(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dataID := types.NewPrimitiveU64(0)
	err = dataID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.PerpetuateObject(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	deleteLastObject := types.NewPrimitiveBool(false)
	err = deleteLastObject.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.PerpetuateObject(fmt.Errorf("Failed to read deleteLastObject from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.PerpetuateObject(nil, packet, callID, persistenceSlotID, dataID, deleteLastObject)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
