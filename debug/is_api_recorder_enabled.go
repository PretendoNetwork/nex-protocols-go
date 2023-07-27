// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// IsAPIRecorderEnabled sets the IsAPIRecorderEnabled handler function
func (protocol *Protocol) IsAPIRecorderEnabled(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.isAPIRecorderEnabledHandler = handler
}

func (protocol *Protocol) handleIsAPIRecorderEnabled(packet nex.PacketInterface) {
	if protocol.isAPIRecorderEnabledHandler == nil {
		globals.Logger.Warning("Debug::IsAPIRecorderEnabled not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.isAPIRecorderEnabledHandler(nil, client, callID)
}
