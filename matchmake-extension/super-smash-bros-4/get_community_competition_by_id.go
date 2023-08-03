// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCommunityCompetitionByID sets the GetCommunityCompetitionByID handler function
func (protocol *Protocol) GetCommunityCompetitionByID(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.getCommunityCompetitionByIDHandler = handler
}

func (protocol *Protocol) handleGetCommunityCompetitionByID(packet nex.PacketInterface) {
	if protocol.getCommunityCompetitionByIDHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetCommunityCompetitionByID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetCommunityCompetitionByID STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getCommunityCompetitionByIDHandler(nil, client, callID, packet.Payload())
}
