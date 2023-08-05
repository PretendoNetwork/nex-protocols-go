// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// EndCommunityCompetitionParticipationByGatheringID sets the EndCommunityCompetitionParticipationByGatheringID handler function
func (protocol *Protocol) EndCommunityCompetitionParticipationByGatheringID(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
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

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.endCommunityCompetitionParticipationByGatheringIDHandler(nil, client, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
