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

	if protocol.RemoveFromBlockList == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::RemoveFromBlockList not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

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
		_, rmcError := protocol.RemoveFromBlockList(fmt.Errorf("Failed to read lstPrincipalID from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RemoveFromBlockList(nil, packet, callID, lstPrincipalID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
