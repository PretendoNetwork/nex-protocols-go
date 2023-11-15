// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRegisterTournamentPlayerInfo(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.RegisterTournamentPlayerInfo == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterTournamentPlayerInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterTournamentPlayerInfo STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.RegisterTournamentPlayerInfo(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
