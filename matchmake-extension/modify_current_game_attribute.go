// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ModifyCurrentGameAttribute sets the ModifyCurrentGameAttribute handler function
func (protocol *Protocol) ModifyCurrentGameAttribute(handler func(err error, client *nex.Client, callID uint32, gid uint32, attribIndex uint32, newValue uint32) uint32) {
	protocol.modifyCurrentGameAttributeHandler = handler
}

func (protocol *Protocol) handleModifyCurrentGameAttribute(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.modifyCurrentGameAttributeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::ModifyCurrentGameAttribute not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.modifyCurrentGameAttributeHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	attribIndex, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.modifyCurrentGameAttributeHandler(fmt.Errorf("Failed to read attribIndex from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	newValue, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.modifyCurrentGameAttributeHandler(fmt.Errorf("Failed to read newValue from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.modifyCurrentGameAttributeHandler(nil, client, callID, gid, attribIndex, newValue)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
