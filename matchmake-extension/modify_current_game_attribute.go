// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleModifyCurrentGameAttribute(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.ModifyCurrentGameAttribute == nil {
		globals.Logger.Warning("MatchmakeExtension::ModifyCurrentGameAttribute not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.ModifyCurrentGameAttribute(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	attribIndex, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.ModifyCurrentGameAttribute(fmt.Errorf("Failed to read attribIndex from parameters. %s", err.Error()), packet, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	newValue, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.ModifyCurrentGameAttribute(fmt.Errorf("Failed to read newValue from parameters. %s", err.Error()), packet, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.ModifyCurrentGameAttribute(nil, packet, callID, gid, attribIndex, newValue)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
