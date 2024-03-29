// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCommunityCompetitionByID sets the GetCommunityCompetitionByID handler function
func (protocol *Protocol) GetCommunityCompetitionByID(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.getCommunityCompetitionByIDHandler = handler
}

func (protocol *Protocol) handleGetCommunityCompetitionByID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getCommunityCompetitionByIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetCommunityCompetitionByID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetCommunityCompetitionByID STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getCommunityCompetitionByIDHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
