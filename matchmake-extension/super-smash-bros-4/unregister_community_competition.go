// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnregisterCommunityCompetition sets the UnregisterCommunityCompetition handler function
func (protocol *Protocol) UnregisterCommunityCompetition(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.unregisterCommunityCompetitionHandler = handler
}

func (protocol *Protocol) handleUnregisterCommunityCompetition(packet nex.PacketInterface) {
	if protocol.unregisterCommunityCompetitionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::UnregisterCommunityCompetition not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::UnregisterCommunityCompetition STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.unregisterCommunityCompetitionHandler(nil, client, callID, packet.Payload())
}
