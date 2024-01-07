// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindCommunityCompetitionsByGatheringID(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.FindCommunityCompetitionsByGatheringID == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::FindCommunityCompetitionsByGatheringID not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::FindCommunityCompetitionsByGatheringID STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.FindCommunityCompetitionsByGatheringID(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
