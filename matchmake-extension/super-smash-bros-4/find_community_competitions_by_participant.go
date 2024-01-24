// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindCommunityCompetitionsByParticipant(packet nex.PacketInterface) {
	if protocol.FindCommunityCompetitionsByParticipant == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::FindCommunityCompetitionsByParticipant not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::FindCommunityCompetitionsByParticipant STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.FindCommunityCompetitionsByParticipant(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
