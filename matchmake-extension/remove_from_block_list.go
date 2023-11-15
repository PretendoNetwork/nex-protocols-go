// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRemoveFromBlockList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.RemoveFromBlockList == nil {
		globals.Logger.Warning("MatchmakeExtension::RemoveFromBlockList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstPrincipalID, err := parametersStream.ReadListPID()
	if err != nil {
		_, errorCode = protocol.RemoveFromBlockList(fmt.Errorf("Failed to read lstPrincipalID from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RemoveFromBlockList(nil, packet, callID, lstPrincipalID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
