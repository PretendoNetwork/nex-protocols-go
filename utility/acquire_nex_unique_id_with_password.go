// Package protocol implements the Utility protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcquireNexUniqueIDWithPassword sets the AcquireNexUniqueIDWithPassword handler function
func (protocol *Protocol) AcquireNexUniqueIDWithPassword(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.acquireNexUniqueIDWithPasswordHandler = handler
}

func (protocol *Protocol) handleAcquireNexUniqueIDWithPassword(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.acquireNexUniqueIDWithPasswordHandler == nil {
		globals.Logger.Warning("Utility::AcquireNexUniqueIDWithPassword not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.acquireNexUniqueIDWithPasswordHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
