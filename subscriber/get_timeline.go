// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetTimeline(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetTimeline == nil {
		globals.Logger.Warning("Subscriber::GetTimeline not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Subscriber::GetTimeline STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.GetTimeline(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
