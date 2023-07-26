// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAPICallSummary sets the GetAPICallSummary handler function
func (protocol *Protocol) GetAPICallSummary(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetAPICallSummaryHandler = handler
}

func (protocol *Protocol) handleGetAPICallSummary(packet nex.PacketInterface) {
	globals.Logger.Warning("Debug::GetAPICallSummary STUBBED")

	if protocol.GetAPICallSummaryHandler == nil {
		globals.Logger.Warning("Debug::GetAPICallSummary not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
}
