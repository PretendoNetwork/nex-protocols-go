// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTournamentProfiles sets the GetTournamentProfiles handler function
func (protocol *Protocol) GetTournamentProfiles(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.getTournamentProfilesHandler = handler
}

func (protocol *Protocol) handleGetTournamentProfiles(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getTournamentProfilesHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentProfiles not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentProfiles STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.getTournamentProfilesHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
