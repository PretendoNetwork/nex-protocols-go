// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// EnableAPIRecorder sets the EnableAPIRecorder handler function
func (protocol *Protocol) EnableAPIRecorder(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.enableAPIRecorderHandler = handler
}

func (protocol *Protocol) handleEnableAPIRecorder(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.enableAPIRecorderHandler == nil {
		globals.Logger.Warning("Debug::EnableAPIRecorder not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.enableAPIRecorderHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
