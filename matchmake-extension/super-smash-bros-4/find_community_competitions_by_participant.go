// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindCommunityCompetitionsByParticipant sets the FindCommunityCompetitionsByParticipant handler function
func (protocol *Protocol) FindCommunityCompetitionsByParticipant(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.findCommunityCompetitionsByParticipantHandler = handler
}

func (protocol *Protocol) handleFindCommunityCompetitionsByParticipant(packet nex.PacketInterface) {
	if protocol.findCommunityCompetitionsByParticipantHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::FindCommunityCompetitionsByParticipant not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::FindCommunityCompetitionsByParticipant STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.findCommunityCompetitionsByParticipantHandler(nil, client, callID, packet.Payload())
}
