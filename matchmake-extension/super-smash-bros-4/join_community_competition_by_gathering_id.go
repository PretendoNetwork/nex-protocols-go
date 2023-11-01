// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// JoinCommunityCompetitionByGatheringID sets the JoinCommunityCompetitionByGatheringID handler function
func (protocol *Protocol) JoinCommunityCompetitionByGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.joinCommunityCompetitionByGatheringIDHandler = handler
}

func (protocol *Protocol) handleJoinCommunityCompetitionByGatheringID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.joinCommunityCompetitionByGatheringIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinCommunityCompetitionByGatheringID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinCommunityCompetitionByGatheringID STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.joinCommunityCompetitionByGatheringIDHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
