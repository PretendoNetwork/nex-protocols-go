// Package protocol implements the Health protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FixSanityErrors sets the FixSanityErrors handler function
func (protocol *Protocol) FixSanityErrors(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.fixSanityErrorsHandler = handler
}

func (protocol *Protocol) handleFixSanityErrors(packet nex.PacketInterface) {
	if protocol.fixSanityErrorsHandler == nil {
		globals.Logger.Warning("Health::FixSanityErrors not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.fixSanityErrorsHandler(nil, client, callID)
}
