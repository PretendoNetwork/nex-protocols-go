// Package storage_manager implements the StorageManager NEX protocol
package storage_manager

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ActivateWithCardID sets the ActivateWithCardID handler function
func (protocol *StorageManagerProtocol) ActivateWithCardID(handler func(err error, client *nex.Client, callID uint32, unknown uint8, cardID uint64)) {
	protocol.activateWithCardIDHandler = handler
}

func (protocol *StorageManagerProtocol) handleActivateWithCardID(packet nex.PacketInterface) {
	if protocol.activateWithCardIDHandler == nil {
		globals.Logger.Warning("StorageManager::ActivateWithCardID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown, err := parametersStream.ReadUInt8()
	if err != nil {
		go protocol.activateWithCardIDHandler(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	cardID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.activateWithCardIDHandler(fmt.Errorf("Failed to read cardID from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.activateWithCardIDHandler(nil, client, callID, unknown, cardID)
}
