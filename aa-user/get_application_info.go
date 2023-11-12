// Package protocol implements the AAUser protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetApplicationInfo sets the GetApplicationInfo handler function
func (protocol *Protocol) GetApplicationInfo(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.getApplicationInfoHandler = handler
}

func (protocol *Protocol) handleGetApplicationInfo(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getApplicationInfoHandler == nil {
		globals.Logger.Warning("AAUser::GetApplicationInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.getApplicationInfoHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
