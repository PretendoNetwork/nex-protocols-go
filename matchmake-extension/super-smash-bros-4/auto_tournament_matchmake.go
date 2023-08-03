// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AutoTournamentMatchmake sets the AutoTournamentMatchmake handler function
func (protocol *Protocol) AutoTournamentMatchmake(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.autoTournamentMatchmakeHandler = handler
}

func (protocol *Protocol) handleAutoTournamentMatchmake(packet nex.PacketInterface) {
	if protocol.autoTournamentMatchmakeHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::AutoTournamentMatchmake not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::AutoTournamentMatchmake STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.autoTournamentMatchmakeHandler(nil, client, callID, packet.Payload())
}
