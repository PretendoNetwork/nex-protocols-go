// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetAPICallSummary(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetAPICallSummary == nil {
		globals.Logger.Warning("Debug::GetAPICallSummary not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Debug::GetAPICallSummary STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	rmcMessage, errorCode := protocol.GetAPICallSummary(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
