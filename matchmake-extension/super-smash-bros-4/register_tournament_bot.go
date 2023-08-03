// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterTournamentBot sets the RegisterTournamentBot handler function
func (protocol *Protocol) RegisterTournamentBot(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.registerTournamentBotHandler = handler
}

func (protocol *Protocol) handleRegisterTournamentBot(packet nex.PacketInterface) {
	if protocol.registerTournamentBotHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterTournamentBot not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterTournamentBot STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.registerTournamentBotHandler(nil, client, callID, packet.Payload())
}
