// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindCommunityCompetitionsByGatheringID sets the FindCommunityCompetitionsByGatheringID handler function
func (protocol *Protocol) FindCommunityCompetitionsByGatheringID(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.findCommunityCompetitionsByGatheringIDHandler = handler
}

func (protocol *Protocol) handleFindCommunityCompetitionsByGatheringID(packet nex.PacketInterface) {
	if protocol.findCommunityCompetitionsByGatheringIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::FindCommunityCompetitionsByGatheringID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::FindCommunityCompetitionsByGatheringID STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.findCommunityCompetitionsByGatheringIDHandler(nil, client, callID, packet.Payload())
}
