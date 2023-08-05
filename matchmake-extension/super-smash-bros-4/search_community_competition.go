// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchCommunityCompetition sets the SearchCommunityCompetition handler function
func (protocol *Protocol) SearchCommunityCompetition(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.searchCommunityCompetitionHandler = handler
}

func (protocol *Protocol) handleSearchCommunityCompetition(packet nex.PacketInterface) {
	if protocol.searchCommunityCompetitionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SearchCommunityCompetition not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SearchCommunityCompetition STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.searchCommunityCompetitionHandler(nil, client, callID, packet.Payload())
}
