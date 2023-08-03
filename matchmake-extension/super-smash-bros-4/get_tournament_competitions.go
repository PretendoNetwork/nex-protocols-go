// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTournamentCompetitions sets the GetTournamentCompetitions handler function
func (protocol *Protocol) GetTournamentCompetitions(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.getTournamentCompetitionsHandler = handler
}

func (protocol *Protocol) handleGetTournamentCompetitions(packet nex.PacketInterface) {
	if protocol.getTournamentCompetitionsHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentCompetitions not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentCompetitions STUBBED")


	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getTournamentCompetitionsHandler(nil, client, callID, packet.Payload())
}