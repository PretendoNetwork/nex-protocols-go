// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUnperpetuateObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UnperpetuateObject == nil {
		globals.Logger.Warning("DataStore::UnperpetuateObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	persistenceSlotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		errorCode = protocol.UnperpetuateObject(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	deleteLastObject, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.UnperpetuateObject(fmt.Errorf("Failed to read deleteLastObject from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.UnperpetuateObject(nil, packet, callID, persistenceSlotID, deleteLastObject)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
