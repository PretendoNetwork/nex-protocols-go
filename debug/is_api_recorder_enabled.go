package debug

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// IsAPIRecorderEnabled sets the IsAPIRecorderEnabled handler function
func (protocol *DebugProtocol) IsAPIRecorderEnabled(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.IsAPIRecorderEnabledHandler = handler
}

func (protocol *DebugProtocol) handleIsAPIRecorderEnabled(packet nex.PacketInterface) {
	if protocol.IsAPIRecorderEnabledHandler == nil {
		globals.Logger.Warning("Debug::IsAPIRecorderEnabled not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.IsAPIRecorderEnabledHandler(nil, client, callID)
}
