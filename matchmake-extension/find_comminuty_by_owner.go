// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindCommunityByOwner sets the FindCommunityByOwner handler function
func (protocol *MatchmakeExtensionProtocol) FindCommunityByOwner(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.findCommunityByOwnerHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleFindCommunityByOwner(packet nex.PacketInterface) {
	if protocol.findCommunityByOwnerHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindCommunityByOwner not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	go protocol.findCommunityByOwnerHandler(nil, client, callID)
}
