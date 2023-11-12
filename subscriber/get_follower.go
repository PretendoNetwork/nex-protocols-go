// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFollower sets the GetFollower handler function
func (protocol *Protocol) GetFollower(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.getFollowerHandler = handler
}

func (protocol *Protocol) handleGetFollower(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getFollowerHandler == nil {
		globals.Logger.Warning("Subscriber::GetFollower not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Subscriber::GetFollower STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.getFollowerHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
