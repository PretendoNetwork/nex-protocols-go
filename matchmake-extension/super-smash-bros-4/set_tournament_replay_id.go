// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSetTournamentReplayID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.SetTournamentReplayID == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SetTournamentReplayID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SetTournamentReplayID STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.SetTournamentReplayID(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
