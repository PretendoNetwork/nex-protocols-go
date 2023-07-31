// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindCommunityByOwner sets the FindCommunityByOwner handler function
func (protocol *Protocol) FindCommunityByOwner(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.findCommunityByOwnerHandler = handler
}

func (protocol *Protocol) handleFindCommunityByOwner(packet nex.PacketInterface) {
	if protocol.findCommunityByOwnerHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindCommunityByOwner not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtension::FindCommunityByOwner STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	go protocol.findCommunityByOwnerHandler(nil, client, callID, packet.Payload())
}
