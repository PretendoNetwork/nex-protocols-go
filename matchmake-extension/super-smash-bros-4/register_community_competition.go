// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRegisterCommunityCompetition(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.RegisterCommunityCompetition == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterCommunityCompetition not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterCommunityCompetition STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.RegisterCommunityCompetition(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
