// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DisableAPIRecorder sets the DisableAPIRecorder handler function
func (protocol *Protocol) DisableAPIRecorder(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.DisableAPIRecorderHandler = handler
}

func (protocol *Protocol) handleDisableAPIRecorder(packet nex.PacketInterface) {
	if protocol.DisableAPIRecorderHandler == nil {
		globals.Logger.Warning("Debug::DisableAPIRecorder not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.DisableAPIRecorderHandler(nil, client, callID)
}
