package debug

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetApiCallSummary sets the GetApiCallSummary handler function
func (protocol *DebugProtocol) GetApiCallSummary(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetApiCallSummaryHandler = handler
}

func (protocol *DebugProtocol) HandleGetApiCallSummary(packet nex.PacketInterface) {
	globals.Logger.Warning("Debug::GetApiCallSummary STUBBED")

	if protocol.GetApiCallSummaryHandler == nil {
		globals.Logger.Warning("Debug::GetApiCallSummary not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
}
