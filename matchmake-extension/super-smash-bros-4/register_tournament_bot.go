// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterTournamentBot sets the RegisterTournamentBot handler function
func (protocol *Protocol) RegisterTournamentBot(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.registerTournamentBotHandler = handler
}

func (protocol *Protocol) handleRegisterTournamentBot(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.registerTournamentBotHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterTournamentBot not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterTournamentBot STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.registerTournamentBotHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
