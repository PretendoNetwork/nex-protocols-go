// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleJoinOrCreateMatchmakeSession(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.JoinOrCreateMatchmakeSession == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinOrCreateMatchmakeSession not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinOrCreateMatchmakeSession STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.JoinOrCreateMatchmakeSession(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
