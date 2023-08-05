// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterTournamentPlayerInfo sets the RegisterTournamentPlayerInfo handler function
func (protocol *Protocol) RegisterTournamentPlayerInfo(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.registerTournamentPlayerInfoHandler = handler
}

func (protocol *Protocol) handleRegisterTournamentPlayerInfo(packet nex.PacketInterface) {
	if protocol.registerTournamentPlayerInfoHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterTournamentPlayerInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::RegisterTournamentPlayerInfo STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.registerTournamentPlayerInfoHandler(nil, client, callID, packet.Payload())
}
