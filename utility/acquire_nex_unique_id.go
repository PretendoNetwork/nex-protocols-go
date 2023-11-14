// Package protocol implements the Utility protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAcquireNexUniqueID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AcquireNexUniqueID == nil {
		globals.Logger.Warning("Utility::AcquireNexUniqueID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.AcquireNexUniqueID(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
