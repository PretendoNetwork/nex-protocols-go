// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAPICallSummary sets the GetAPICallSummary handler function
func (protocol *Protocol) GetAPICallSummary(handler func(err error, packet nex.PacketInterface, callID uint32, pakcetPayload []byte) uint32) {
	protocol.getAPICallSummaryHandler = handler
}

func (protocol *Protocol) handleGetAPICallSummary(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getAPICallSummaryHandler == nil {
		globals.Logger.Warning("Debug::GetAPICallSummary not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Debug::GetAPICallSummary STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	errorCode = protocol.getAPICallSummaryHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
