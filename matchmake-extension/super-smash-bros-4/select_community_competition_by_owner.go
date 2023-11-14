// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSelectCommunityCompetitionByOwner(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.SelectCommunityCompetitionByOwner == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SelectCommunityCompetitionByOwner not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SelectCommunityCompetitionByOwner STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.SelectCommunityCompetitionByOwner(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
