// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DebugJoinCommunityCompetition sets the DebugJoinCommunityCompetition handler function
func (protocol *Protocol) DebugJoinCommunityCompetition(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.debugJoinCommunityCompetitionHandler = handler
}

func (protocol *Protocol) handleDebugJoinCommunityCompetition(packet nex.PacketInterface) {
	if protocol.debugJoinCommunityCompetitionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugJoinCommunityCompetition not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugJoinCommunityCompetition STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.debugJoinCommunityCompetitionHandler(nil, client, callID, packet.Payload())
}
