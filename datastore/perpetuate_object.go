// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PerpetuateObject sets the PerpetuateObject handler function
func (protocol *Protocol) PerpetuateObject(handler func(err error, client *nex.Client, callID uint32, persistenceSlotID uint16, dataID uint64, deleteLastObject bool)) {
	protocol.perpetuateObjectHandler = handler
}

func (protocol *Protocol) handlePerpetuateObject(packet nex.PacketInterface) {
	if protocol.perpetuateObjectHandler == nil {
		globals.Logger.Warning("DataStore::PerpetuateObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	persistenceSlotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		go protocol.perpetuateObjectHandler(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), client, callID, 0, 0, false)
		return
	}

	dataID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.perpetuateObjectHandler(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), client, callID, 0, 0, false)
		return
	}

	deleteLastObject, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.perpetuateObjectHandler(fmt.Errorf("Failed to read deleteLastObject from parameters. %s", err.Error()), client, callID, 0, 0, false)
		return
	}

	go protocol.perpetuateObjectHandler(nil, client, callID, persistenceSlotID, dataID, deleteLastObject)
}