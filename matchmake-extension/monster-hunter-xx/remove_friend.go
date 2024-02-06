// Package protocol implements the MatchmakeExtensionMonsterHunterXX protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRemoveFriend(packet nex.PacketInterface) {
	if protocol.RemoveFriend == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtensionMonsterHunterXX::RemoveFriend not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	pid := types.NewPID(0)

	err := pid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RemoveFriend(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RemoveFriend(nil, packet, callID, pid)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
