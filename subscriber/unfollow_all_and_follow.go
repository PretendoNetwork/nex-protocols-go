// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnfollowAllAndFollow sets the UnfollowAllAndFollow handler function
func (protocol *Protocol) UnfollowAllAndFollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.unfollowAllAndFollowHandler = handler
}

func (protocol *Protocol) handleUnfollowAllAndFollow(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.unfollowAllAndFollowHandler == nil {
		globals.Logger.Warning("Subscriber::UnfollowAllAndFollow not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Subscriber::UnfollowAllAndFollow STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.unfollowAllAndFollowHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
