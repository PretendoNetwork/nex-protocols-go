// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostCommunityCompetitionMatchResult sets the PostCommunityCompetitionMatchResult handler function
func (protocol *Protocol) PostCommunityCompetitionMatchResult(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.postCommunityCompetitionMatchResultHandler = handler
}

func (protocol *Protocol) handlePostCommunityCompetitionMatchResult(packet nex.PacketInterface) {
	if protocol.postCommunityCompetitionMatchResultHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::PostCommunityCompetitionMatchResult not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::PostCommunityCompetitionMatchResult STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.postCommunityCompetitionMatchResultHandler(nil, client, callID, packet.Payload())
}
