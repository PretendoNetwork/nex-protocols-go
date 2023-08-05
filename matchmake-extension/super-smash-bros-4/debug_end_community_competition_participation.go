// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DebugEndCommunityCompetitionParticipation sets the DebugEndCommunityCompetitionParticipation handler function
func (protocol *Protocol) DebugEndCommunityCompetitionParticipation(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.debugEndCommunityCompetitionParticipationHandler = handler
}

func (protocol *Protocol) handleDebugEndCommunityCompetitionParticipation(packet nex.PacketInterface) {
	if protocol.debugEndCommunityCompetitionParticipationHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugEndCommunityCompetitionParticipation not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugEndCommunityCompetitionParticipation STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.debugEndCommunityCompetitionParticipationHandler(nil, client, callID, packet.Payload())
}
