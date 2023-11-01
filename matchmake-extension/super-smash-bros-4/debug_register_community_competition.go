// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DebugRegisterCommunityCompetition sets the DebugRegisterCommunityCompetition handler function
func (protocol *Protocol) DebugRegisterCommunityCompetition(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.debugRegisterCommunityCompetitionHandler = handler
}

func (protocol *Protocol) handleDebugRegisterCommunityCompetition(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.debugRegisterCommunityCompetitionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugRegisterCommunityCompetition not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugRegisterCommunityCompetition STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.debugRegisterCommunityCompetitionHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
