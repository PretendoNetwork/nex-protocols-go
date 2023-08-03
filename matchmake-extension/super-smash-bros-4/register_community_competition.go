// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterCommunityCompetition sets the RegisterCommunityCompetition handler function
func (protocol *Protocol) RegisterCommunityCompetition(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.registerCommunityCompetitionHandler = handler
}

func (protocol *Protocol) handleRegisterCommunityCompetition(packet nex.PacketInterface) {
	if protocol.registerCommunityCompetitionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterCommunityCompetition not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterCommunityCompetition STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.registerCommunityCompetitionHandler(nil, client, callID, packet.Payload())
}
