package debug

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// IsApiRecorderEnabled sets the IsApiRecorderEnabled handler function
func (protocol *DebugProtocol) IsApiRecorderEnabled(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.IsApiRecorderEnabledHandler = handler
}

func (protocol *DebugProtocol) HandleIsApiRecorderEnabled(packet nex.PacketInterface) {
	if protocol.IsApiRecorderEnabledHandler == nil {
		globals.Logger.Warning("Debug::IsApiRecorderEnabled not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.IsApiRecorderEnabledHandler(nil, client, callID)
}
