// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTournamentReplayIDs sets the GetTournamentReplayIDs handler function
func (protocol *Protocol) GetTournamentReplayIDs(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.getTournamentReplayIDsHandler = handler
}

func (protocol *Protocol) handleGetTournamentReplayIDs(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getTournamentReplayIDsHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentReplayIDs not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentReplayIDs STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.getTournamentReplayIDsHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
