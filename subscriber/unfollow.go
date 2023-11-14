// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUnfollow(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.Unfollow == nil {
		globals.Logger.Warning("Subscriber::Unfollow not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Subscriber::Unfollow STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.Unfollow(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
