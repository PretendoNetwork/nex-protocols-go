// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// JoinCommunityCompetition sets the JoinCommunityCompetition handler function
func (protocol *Protocol) JoinCommunityCompetition(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.joinCommunityCompetitionHandler = handler
}

func (protocol *Protocol) handleJoinCommunityCompetition(packet nex.PacketInterface) {
	if protocol.joinCommunityCompetitionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinCommunityCompetition not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinCommunityCompetition STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.joinCommunityCompetitionHandler(nil, client, callID, packet.Payload())
}
