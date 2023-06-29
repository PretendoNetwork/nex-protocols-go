package debug

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DisableApiRecorder sets the DisableApiRecorder handler function
func (protocol *DebugProtocol) DisableApiRecorder(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.DisableApiRecorderHandler = handler
}

func (protocol *DebugProtocol) handleDisableApiRecorder(packet nex.PacketInterface) {
	if protocol.DisableApiRecorderHandler == nil {
		globals.Logger.Warning("Debug::DisableApiRecorder not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.DisableApiRecorderHandler(nil, client, callID)
}
