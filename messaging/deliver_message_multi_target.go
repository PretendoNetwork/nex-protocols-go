// Package protocol implements the Messaging protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeliverMessageMultiTarget(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.DeliverMessageMultiTarget == nil {
		globals.Logger.Warning("Messaging::DeliverMessageMultiTarget not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Messaging::DeliverMessageMultiTarget STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	rmcMessage, errorCode := protocol.DeliverMessageMultiTarget(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
