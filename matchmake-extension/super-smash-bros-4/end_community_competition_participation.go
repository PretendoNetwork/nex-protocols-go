// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// EndCommunityCompetitionParticipation sets the EndCommunityCompetitionParticipation handler function
func (protocol *Protocol) EndCommunityCompetitionParticipation(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.endCommunityCompetitionParticipationHandler = handler
}

func (protocol *Protocol) handleEndCommunityCompetitionParticipation(packet nex.PacketInterface) {
	if protocol.endCommunityCompetitionParticipationHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::EndCommunityCompetitionParticipation not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::EndCommunityCompetitionParticipation STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.endCommunityCompetitionParticipationHandler(nil, client, callID, packet.Payload())
}
