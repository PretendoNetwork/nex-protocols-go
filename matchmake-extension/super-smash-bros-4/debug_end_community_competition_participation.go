// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DebugEndCommunityCompetitionParticipation sets the DebugEndCommunityCompetitionParticipation handler function
func (protocol *Protocol) DebugEndCommunityCompetitionParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.debugEndCommunityCompetitionParticipationHandler = handler
}

func (protocol *Protocol) handleDebugEndCommunityCompetitionParticipation(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.debugEndCommunityCompetitionParticipationHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugEndCommunityCompetitionParticipation not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugEndCommunityCompetitionParticipation STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.debugEndCommunityCompetitionParticipationHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
