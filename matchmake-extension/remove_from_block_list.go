// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRemoveFromBlockList(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.RemoveFromBlockList == nil {
		globals.Logger.Warning("MatchmakeExtension::RemoveFromBlockList not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lstPrincipalID := types.NewList[*types.PID]()
	lstPrincipalID.Type = types.NewPID(0)
	err = lstPrincipalID.ExtractFrom(parametersStream)
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
