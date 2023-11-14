// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleJoinCommunityCompetitionByGatheringID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.JoinCommunityCompetitionByGatheringID == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinCommunityCompetitionByGatheringID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinCommunityCompetitionByGatheringID STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.JoinCommunityCompetitionByGatheringID(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
