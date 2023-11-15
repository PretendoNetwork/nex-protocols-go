// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handlePerpetuateObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.PerpetuateObject == nil {
		globals.Logger.Warning("DataStore::PerpetuateObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	persistenceSlotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		_, errorCode = protocol.PerpetuateObject(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), packet, callID, 0, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dataID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		_, errorCode = protocol.PerpetuateObject(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), packet, callID, 0, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	deleteLastObject, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.PerpetuateObject(fmt.Errorf("Failed to read deleteLastObject from parameters. %s", err.Error()), packet, callID, 0, 0, false)
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
