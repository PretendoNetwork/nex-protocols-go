// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCommunityCompetitionRanking sets the GetCommunityCompetitionRanking handler function
func (protocol *Protocol) GetCommunityCompetitionRanking(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.getCommunityCompetitionRankingHandler = handler
}

func (protocol *Protocol) handleGetCommunityCompetitionRanking(packet nex.PacketInterface) {
	if protocol.getCommunityCompetitionRankingHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetCommunityCompetitionRanking not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetCommunityCompetitionRanking STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getCommunityCompetitionRankingHandler(nil, client, callID, packet.Payload())
}
