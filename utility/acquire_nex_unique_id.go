// Package protocol implements the Utility protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcquireNexUniqueID sets the AcquireNexUniqueID handler function
func (protocol *Protocol) AcquireNexUniqueID(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.acquireNexUniqueIDHandler = handler
}

func (protocol *Protocol) handleAcquireNexUniqueID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.acquireNexUniqueIDHandler == nil {
		globals.Logger.Warning("Utility::AcquireNexUniqueID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.acquireNexUniqueIDHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
