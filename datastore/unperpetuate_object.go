// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnperpetuateObject sets the UnperpetuateObject handler function
func (protocol *Protocol) UnperpetuateObject(handler func(err error, client *nex.Client, callID uint32, persistenceSlotID uint16, deleteLastObject bool) uint32) {
	protocol.unperpetuateObjectHandler = handler
}

func (protocol *Protocol) handleUnperpetuateObject(packet nex.PacketInterface) {
	if protocol.unperpetuateObjectHandler == nil {
		globals.Logger.Warning("DataStore::UnperpetuateObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	persistenceSlotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		go protocol.unperpetuateObjectHandler(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), client, callID, 0, false)
		return
	}

	deleteLastObject, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.unperpetuateObjectHandler(fmt.Errorf("Failed to read deleteLastObject from parameters. %s", err.Error()), client, callID, 0, false)
		return
	}

	go protocol.unperpetuateObjectHandler(nil, client, callID, persistenceSlotID, deleteLastObject)
}
