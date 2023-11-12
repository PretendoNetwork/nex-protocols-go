// Package protocol implements the StorageManager protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ActivateWithCardID sets the ActivateWithCardID handler function
func (protocol *Protocol) ActivateWithCardID(handler func(err error, packet nex.PacketInterface, callID uint32, unknown uint8, cardID uint64) uint32) {
	protocol.activateWithCardIDHandler = handler
}

func (protocol *Protocol) handleActivateWithCardID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.activateWithCardIDHandler == nil {
		globals.Logger.Warning("StorageManager::ActivateWithCardID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown, err := parametersStream.ReadUInt8()
	if err != nil {
		errorCode = protocol.activateWithCardIDHandler(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	cardID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.activateWithCardIDHandler(fmt.Errorf("Failed to read cardID from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.activateWithCardIDHandler(nil, packet, callID, unknown, cardID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
