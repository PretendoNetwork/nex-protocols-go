// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCommunityCompetitions sets the GetCommunityCompetitions handler function
func (protocol *Protocol) GetCommunityCompetitions(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.getCommunityCompetitionsHandler = handler
}

func (protocol *Protocol) handleGetCommunityCompetitions(packet nex.PacketInterface) {
	if protocol.getCommunityCompetitionsHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetCommunityCompetitions not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetCommunityCompetitions STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getCommunityCompetitionsHandler(nil, client, callID, packet.Payload())
}
