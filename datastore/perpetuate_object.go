// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PerpetuateObject sets the PerpetuateObject handler function
func (protocol *Protocol) PerpetuateObject(handler func(err error, packet nex.PacketInterface, callID uint32, persistenceSlotID uint16, dataID uint64, deleteLastObject bool) uint32) {
	protocol.perpetuateObjectHandler = handler
}

func (protocol *Protocol) handlePerpetuateObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.perpetuateObjectHandler == nil {
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
		errorCode = protocol.perpetuateObjectHandler(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), packet, callID, 0, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dataID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.perpetuateObjectHandler(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), packet, callID, 0, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	deleteLastObject, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.perpetuateObjectHandler(fmt.Errorf("Failed to read deleteLastObject from parameters. %s", err.Error()), packet, callID, 0, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.perpetuateObjectHandler(nil, packet, callID, persistenceSlotID, dataID, deleteLastObject)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
