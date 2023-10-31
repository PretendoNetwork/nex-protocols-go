// Package protocol implements the Health protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FixSanityErrors sets the FixSanityErrors handler function
func (protocol *Protocol) FixSanityErrors(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.fixSanityErrorsHandler = handler
}

func (protocol *Protocol) handleFixSanityErrors(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.fixSanityErrorsHandler == nil {
		globals.Logger.Warning("Health::FixSanityErrors not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.fixSanityErrorsHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
