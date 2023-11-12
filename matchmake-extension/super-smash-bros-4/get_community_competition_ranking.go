// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCommunityCompetitionRanking sets the GetCommunityCompetitionRanking handler function
func (protocol *Protocol) GetCommunityCompetitionRanking(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.getCommunityCompetitionRankingHandler = handler
}

func (protocol *Protocol) handleGetCommunityCompetitionRanking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getCommunityCompetitionRankingHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetCommunityCompetitionRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetCommunityCompetitionRanking STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.getCommunityCompetitionRankingHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
