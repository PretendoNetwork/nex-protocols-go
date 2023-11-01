// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTimeline sets the GetTimeline handler function
func (protocol *Protocol) GetTimeline(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.getTimelineHandler = handler
}

func (protocol *Protocol) handleGetTimeline(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getTimelineHandler == nil {
		globals.Logger.Warning("Subscriber::GetTimeline not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Subscriber::GetTimeline STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getTimelineHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
