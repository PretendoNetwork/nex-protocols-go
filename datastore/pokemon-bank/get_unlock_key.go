// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetUnlockKey sets the GetUnlockKey handler function
func (protocol *Protocol) GetUnlockKey(handler func(err error, packet nex.PacketInterface, callID uint32, challengeValue uint32) uint32) {
	protocol.getUnlockKeyHandler = handler
}

func (protocol *Protocol) handleGetUnlockKey(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getUnlockKeyHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::GetUnlockKey not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	challengeValue, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getUnlockKeyHandler(fmt.Errorf("Failed to read challengeValue from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getUnlockKeyHandler(nil, packet, callID, challengeValue)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
