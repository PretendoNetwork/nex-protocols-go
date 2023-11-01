// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFollowing sets the GetFollowing handler function
func (protocol *Protocol) GetFollowing(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.getFollowingHandler = handler
}

func (protocol *Protocol) handleGetFollowing(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getFollowingHandler == nil {
		globals.Logger.Warning("Subscriber::GetFollowing not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Subscriber::GetFollowing STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getFollowingHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
