// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Unfollow sets the Unfollow handler function
func (protocol *Protocol) Unfollow(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.unfollowHandler = handler
}

func (protocol *Protocol) handleUnfollow(packet nex.PacketInterface) {
	if protocol.unfollowHandler == nil {
		globals.Logger.Warning("Subscriber::Unfollow not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Subscriber::Unfollow STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.unfollowHandler(nil, client, callID, packet.Payload())
}
