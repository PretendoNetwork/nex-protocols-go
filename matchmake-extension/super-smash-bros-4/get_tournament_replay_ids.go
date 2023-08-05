// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTournamentReplayIDs sets the GetTournamentReplayIDs handler function
func (protocol *Protocol) GetTournamentReplayIDs(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.getTournamentReplayIDsHandler = handler
}

func (protocol *Protocol) handleGetTournamentReplayIDs(packet nex.PacketInterface) {
	if protocol.getTournamentReplayIDsHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentReplayIDs not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentReplayIDs STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getTournamentReplayIDsHandler(nil, client, callID, packet.Payload())
}
