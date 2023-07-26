// Package protocol implements the Health protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FixSanityErrors sets the FixSanityErrors handler function
func (protocol *Protocol) FixSanityErrors(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.FixSanityErrorsHandler = handler
}

func (protocol *Protocol) handleFixSanityErrors(packet nex.PacketInterface) {
	if protocol.FixSanityErrorsHandler == nil {
		globals.Logger.Warning("Health::FixSanityErrors not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.FixSanityErrorsHandler(nil, client, callID)
}
