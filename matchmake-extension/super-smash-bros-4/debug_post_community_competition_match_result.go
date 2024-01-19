// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDebugPostCommunityCompetitionMatchResult(packet nex.PacketInterface) {
	if protocol.DebugPostCommunityCompetitionMatchResult == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugPostCommunityCompetitionMatchResult not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugPostCommunityCompetitionMatchResult STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.DebugPostCommunityCompetitionMatchResult(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
