// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Unfollow sets the Unfollow handler function
func (protocol *Protocol) Unfollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.unfollowHandler = handler
}

func (protocol *Protocol) handleUnfollow(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.unfollowHandler == nil {
		globals.Logger.Warning("Subscriber::Unfollow not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Subscriber::Unfollow STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.unfollowHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
