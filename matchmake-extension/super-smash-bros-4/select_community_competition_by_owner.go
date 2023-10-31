// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SelectCommunityCompetitionByOwner sets the SelectCommunityCompetitionByOwner handler function
func (protocol *Protocol) SelectCommunityCompetitionByOwner(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.selectCommunityCompetitionByOwnerHandler = handler
}

func (protocol *Protocol) handleSelectCommunityCompetitionByOwner(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.selectCommunityCompetitionByOwnerHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SelectCommunityCompetitionByOwner not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::SelectCommunityCompetitionByOwner STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.selectCommunityCompetitionByOwnerHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
