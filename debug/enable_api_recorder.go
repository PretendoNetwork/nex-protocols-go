// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// EnableAPIRecorder sets the EnableAPIRecorder handler function
func (protocol *Protocol) EnableAPIRecorder(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.enableAPIRecorderHandler = handler
}

func (protocol *Protocol) handleEnableAPIRecorder(packet nex.PacketInterface) {
	if protocol.enableAPIRecorderHandler == nil {
		globals.Logger.Warning("Debug::EnableAPIRecorder not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.enableAPIRecorderHandler(nil, client, callID)
}
