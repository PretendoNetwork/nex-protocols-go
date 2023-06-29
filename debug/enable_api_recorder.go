package debug

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// EnableAPIRecorder sets the EnableAPIRecorder handler function
func (protocol *DebugProtocol) EnableAPIRecorder(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.EnableAPIRecorderHandler = handler
}

func (protocol *DebugProtocol) handleEnableAPIRecorder(packet nex.PacketInterface) {
	if protocol.EnableAPIRecorderHandler == nil {
		globals.Logger.Warning("Debug::EnableAPIRecorder not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.EnableAPIRecorderHandler(nil, client, callID)
}
