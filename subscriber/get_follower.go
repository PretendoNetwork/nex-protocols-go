// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetFollower(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetFollower == nil {
		globals.Logger.Warning("Subscriber::GetFollower not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Subscriber::GetFollower STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.GetFollower(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
