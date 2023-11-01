// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterTournamentPlayerInfo sets the RegisterTournamentPlayerInfo handler function
func (protocol *Protocol) RegisterTournamentPlayerInfo(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.registerTournamentPlayerInfoHandler = handler
}

func (protocol *Protocol) handleRegisterTournamentPlayerInfo(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.registerTournamentPlayerInfoHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterTournamentPlayerInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterTournamentPlayerInfo STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.registerTournamentPlayerInfoHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
