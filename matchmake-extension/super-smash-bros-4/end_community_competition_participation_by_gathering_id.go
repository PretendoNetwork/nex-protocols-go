// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// EndCommunityCompetitionParticipationByGatheringID sets the EndCommunityCompetitionParticipationByGatheringID handler function
func (protocol *Protocol) EndCommunityCompetitionParticipationByGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.endCommunityCompetitionParticipationByGatheringIDHandler = handler
}

func (protocol *Protocol) handleEndCommunityCompetitionParticipationByGatheringID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.endCommunityCompetitionParticipationByGatheringIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::EndCommunityCompetitionParticipationByGatheringID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::EndCommunityCompetitionParticipationByGatheringID STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.endCommunityCompetitionParticipationByGatheringIDHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
