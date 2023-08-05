// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTournamentProfiles sets the GetTournamentProfiles handler function
func (protocol *Protocol) GetTournamentProfiles(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.getTournamentProfilesHandler = handler
}

func (protocol *Protocol) handleGetTournamentProfiles(packet nex.PacketInterface) {
	if protocol.getTournamentProfilesHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentProfiles not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentProfiles STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getTournamentProfilesHandler(nil, client, callID, packet.Payload())
}
