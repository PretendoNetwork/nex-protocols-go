// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteSimpleSearchObject sets the DeleteSimpleSearchObject handler function
func (protocol *Protocol) DeleteSimpleSearchObject(handler func(err error, packet nex.PacketInterface, callID uint32, objectID uint32) uint32) {
	protocol.deleteSimpleSearchObjectHandler = handler
}

func (protocol *Protocol) handleDeleteSimpleSearchObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deleteSimpleSearchObjectHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::DeleteSimpleSearchObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	objectID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.deleteSimpleSearchObjectHandler(fmt.Errorf("Failed to read objectID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.deleteSimpleSearchObjectHandler(nil, packet, callID, objectID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
