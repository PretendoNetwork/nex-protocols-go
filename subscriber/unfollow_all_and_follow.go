// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnfollowAllAndFollow sets the UnfollowAllAndFollow handler function
func (protocol *Protocol) UnfollowAllAndFollow(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.unfollowAllAndFollowHandler = handler
}

func (protocol *Protocol) handleUnfollowAllAndFollow(packet nex.PacketInterface) {
	if protocol.unfollowAllAndFollowHandler == nil {
		globals.Logger.Warning("Subscriber::UnfollowAllAndFollow not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("Subscriber::UnfollowAllAndFollow STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.unfollowAllAndFollowHandler(nil, client, callID, packet.Payload())
}
