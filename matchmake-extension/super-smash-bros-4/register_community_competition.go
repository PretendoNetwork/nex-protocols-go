// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterCommunityCompetition sets the RegisterCommunityCompetition handler function
func (protocol *Protocol) RegisterCommunityCompetition(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.registerCommunityCompetitionHandler = handler
}

func (protocol *Protocol) handleRegisterCommunityCompetition(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.registerCommunityCompetitionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterCommunityCompetition not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterCommunityCompetition STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.registerCommunityCompetitionHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
