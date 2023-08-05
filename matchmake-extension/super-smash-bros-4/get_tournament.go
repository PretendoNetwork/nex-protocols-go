// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTournament sets the GetTournament handler function
func (protocol *Protocol) GetTournament(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.getTournamentHandler = handler
}

func (protocol *Protocol) handleGetTournament(packet nex.PacketInterface) {
	if protocol.getTournamentHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournament not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournament STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getTournamentHandler(nil, client, callID, packet.Payload())
}
