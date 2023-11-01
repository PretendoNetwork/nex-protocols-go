// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCommunityCompetitions sets the GetCommunityCompetitions handler function
func (protocol *Protocol) GetCommunityCompetitions(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.getCommunityCompetitionsHandler = handler
}

func (protocol *Protocol) handleGetCommunityCompetitions(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getCommunityCompetitionsHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetCommunityCompetitions not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetCommunityCompetitions STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getCommunityCompetitionsHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
