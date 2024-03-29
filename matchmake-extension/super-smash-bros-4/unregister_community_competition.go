// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnregisterCommunityCompetition sets the UnregisterCommunityCompetition handler function
func (protocol *Protocol) UnregisterCommunityCompetition(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.unregisterCommunityCompetitionHandler = handler
}

func (protocol *Protocol) handleUnregisterCommunityCompetition(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.unregisterCommunityCompetitionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::UnregisterCommunityCompetition not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::UnregisterCommunityCompetition STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.unregisterCommunityCompetitionHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
