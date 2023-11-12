// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// StartTournament sets the StartTournament handler function
func (protocol *Protocol) StartTournament(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.startTournamentHandler = handler
}

func (protocol *Protocol) handleStartTournament(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.startTournamentHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::StartTournament not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::StartTournament STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.startTournamentHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
