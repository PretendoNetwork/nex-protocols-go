// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DebugUnregisterCommunityCompetition sets the DebugUnregisterCommunityCompetition handler function
func (protocol *Protocol) DebugUnregisterCommunityCompetition(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.debugUnregisterCommunityCompetitionHandler = handler
}

func (protocol *Protocol) handleDebugUnregisterCommunityCompetition(packet nex.PacketInterface) {
	if protocol.debugUnregisterCommunityCompetitionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugUnregisterCommunityCompetition not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::DebugUnregisterCommunityCompetition STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.debugUnregisterCommunityCompetitionHandler(nil, client, callID, packet.Payload())
}
