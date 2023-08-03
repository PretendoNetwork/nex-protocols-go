// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// StartTournament sets the StartTournament handler function
func (protocol *Protocol) StartTournament(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.startTournamentHandler = handler
}

func (protocol *Protocol) handleStartTournament(packet nex.PacketInterface) {
	if protocol.startTournamentHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::StartTournament not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::StartTournament STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.startTournamentHandler(nil, client, callID, packet.Payload())
}
