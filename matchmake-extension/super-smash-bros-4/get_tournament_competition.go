// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTournamentCompetition sets the GetTournamentCompetition handler function
func (protocol *Protocol) GetTournamentCompetition(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.getTournamentCompetitionHandler = handler
}

func (protocol *Protocol) handleGetTournamentCompetition(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getTournamentCompetitionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentCompetition not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentCompetition STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getTournamentCompetitionHandler(nil, client, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
