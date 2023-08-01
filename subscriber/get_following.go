// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFollowing sets the GetFollowing handler function
func (protocol *Protocol) GetFollowing(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.getFollowingHandler = handler
}

func (protocol *Protocol) handleGetFollowing(packet nex.PacketInterface) {
	if protocol.getFollowingHandler == nil {
		globals.Logger.Warning("Subscriber::GetFollowing not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("Subscriber::GetFollowing STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getFollowingHandler(nil, client, callID, packet.Payload())
}
