// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DebugPostCommunityCompetitionMatchResult sets the DebugPostCommunityCompetitionMatchResult handler function
func (protocol *Protocol) DebugPostCommunityCompetitionMatchResult(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.debugPostCommunityCompetitionMatchResultHandler = handler
}

func (protocol *Protocol) handleDebugPostCommunityCompetitionMatchResult(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.debugPostCommunityCompetitionMatchResultHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugPostCommunityCompetitionMatchResult not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugPostCommunityCompetitionMatchResult STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.debugPostCommunityCompetitionMatchResultHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
