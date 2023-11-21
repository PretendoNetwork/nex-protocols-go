// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRemoveFriendByPrincipalID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.RemoveFriendByPrincipalID == nil {
		globals.Logger.Warning("Friends3DS::RemoveFriendByPrincipalID not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadPID()
	if err != nil {
		_, errorCode = protocol.RemoveFriendByPrincipalID(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RemoveFriendByPrincipalID(nil, packet, callID, pid)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
