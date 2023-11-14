// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetNumFollowers(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetNumFollowers == nil {
		globals.Logger.Warning("Subscriber::GetNumFollowers not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Subscriber::GetNumFollowers STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.GetNumFollowers(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
