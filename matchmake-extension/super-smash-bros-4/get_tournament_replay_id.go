// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTournamentReplayID sets the GetTournamentReplayID handler function
func (protocol *Protocol) GetTournamentReplayID(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.getTournamentReplayIDHandler = handler
}

func (protocol *Protocol) handleGetTournamentReplayID(packet nex.PacketInterface) {
	if protocol.getTournamentReplayIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentReplayID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentReplayID STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getTournamentReplayIDHandler(nil, client, callID, packet.Payload())
}
