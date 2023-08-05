// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnregisterCommunityCompetitionByID sets the UnregisterCommunityCompetitionByID handler function
func (protocol *Protocol) UnregisterCommunityCompetitionByID(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.unregisterCommunityCompetitionByIDHandler = handler
}

func (protocol *Protocol) handleUnregisterCommunityCompetitionByID(packet nex.PacketInterface) {
	if protocol.unregisterCommunityCompetitionByIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::UnregisterCommunityCompetitionByID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::UnregisterCommunityCompetitionByID STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.unregisterCommunityCompetitionByIDHandler(nil, client, callID, packet.Payload())
}
