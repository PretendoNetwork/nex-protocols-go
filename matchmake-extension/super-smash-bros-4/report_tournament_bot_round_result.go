// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportTournamentBotRoundResult sets the ReportTournamentBotRoundResult handler function
func (protocol *Protocol) ReportTournamentBotRoundResult(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.reportTournamentBotRoundResultHandler = handler
}

func (protocol *Protocol) handleReportTournamentBotRoundResult(packet nex.PacketInterface) {
	if protocol.reportTournamentBotRoundResultHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::ReportTournamentBotRoundResult not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::ReportTournamentBotRoundResult STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.reportTournamentBotRoundResultHandler(nil, client, callID, packet.Payload())
}
