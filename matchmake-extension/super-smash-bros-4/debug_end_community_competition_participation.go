// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDebugEndCommunityCompetitionParticipation(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.DebugEndCommunityCompetitionParticipation == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugEndCommunityCompetitionParticipation not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugEndCommunityCompetitionParticipation STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.DebugEndCommunityCompetitionParticipation(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
