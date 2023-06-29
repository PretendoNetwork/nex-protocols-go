package debug

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// EnableApiRecorder sets the EnableApiRecorder handler function
func (protocol *DebugProtocol) EnableApiRecorder(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.EnableApiRecorderHandler = handler
}

func (protocol *DebugProtocol) handleEnableApiRecorder(packet nex.PacketInterface) {
	if protocol.EnableApiRecorderHandler == nil {
		globals.Logger.Warning("Debug::EnableApiRecorder not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.EnableApiRecorderHandler(nil, client, callID)
}
