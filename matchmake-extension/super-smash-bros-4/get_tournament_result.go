// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTournamentResult sets the GetTournamentResult handler function
func (protocol *Protocol) GetTournamentResult(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.getTournamentResultHandler = handler
}

func (protocol *Protocol) handleGetTournamentResult(packet nex.PacketInterface) {
	if protocol.getTournamentResultHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentResult not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentResult STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getTournamentResultHandler(nil, client, callID, packet.Payload())
}
