// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleIsAPIRecorderEnabled(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.IsAPIRecorderEnabled == nil {
		globals.Logger.Warning("Debug::IsAPIRecorderEnabled not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.IsAPIRecorderEnabled(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
