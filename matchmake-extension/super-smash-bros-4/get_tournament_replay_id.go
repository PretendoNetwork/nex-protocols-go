// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTournamentReplayID sets the GetTournamentReplayID handler function
func (protocol *Protocol) GetTournamentReplayID(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.getTournamentReplayIDHandler = handler
}

func (protocol *Protocol) handleGetTournamentReplayID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getTournamentReplayIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentReplayID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentReplayID STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.getTournamentReplayIDHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
