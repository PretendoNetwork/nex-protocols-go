// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SetTournamentReplayID sets the SetTournamentReplayID handler function
func (protocol *Protocol) SetTournamentReplayID(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.setTournamentReplayIDHandler = handler
}

func (protocol *Protocol) handleSetTournamentReplayID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.setTournamentReplayIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SetTournamentReplayID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SetTournamentReplayID STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.setTournamentReplayIDHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
