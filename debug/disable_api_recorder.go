// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DisableAPIRecorder sets the DisableAPIRecorder handler function
func (protocol *Protocol) DisableAPIRecorder(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.disableAPIRecorderHandler = handler
}

func (protocol *Protocol) handleDisableAPIRecorder(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.disableAPIRecorderHandler == nil {
		globals.Logger.Warning("Debug::DisableAPIRecorder not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.disableAPIRecorderHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
