// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DebugJoinCommunityCompetition sets the DebugJoinCommunityCompetition handler function
func (protocol *Protocol) DebugJoinCommunityCompetition(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.debugJoinCommunityCompetitionHandler = handler
}

func (protocol *Protocol) handleDebugJoinCommunityCompetition(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.debugJoinCommunityCompetitionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugJoinCommunityCompetition not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugJoinCommunityCompetition STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.debugJoinCommunityCompetitionHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
