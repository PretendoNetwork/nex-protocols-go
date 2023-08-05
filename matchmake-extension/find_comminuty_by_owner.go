// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindCommunityByOwner sets the FindCommunityByOwner handler function
func (protocol *Protocol) FindCommunityByOwner(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.findCommunityByOwnerHandler = handler
}

func (protocol *Protocol) handleFindCommunityByOwner(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findCommunityByOwnerHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindCommunityByOwner not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtension::FindCommunityByOwner STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	errorCode = protocol.findCommunityByOwnerHandler(nil, client, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
