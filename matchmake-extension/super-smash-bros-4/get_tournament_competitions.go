// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTournamentCompetitions sets the GetTournamentCompetitions handler function
func (protocol *Protocol) GetTournamentCompetitions(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.getTournamentCompetitionsHandler = handler
}

func (protocol *Protocol) handleGetTournamentCompetitions(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getTournamentCompetitionsHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentCompetitions not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentCompetitions STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.getTournamentCompetitionsHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
