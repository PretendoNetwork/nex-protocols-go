// Package protocol implements the Ticket Granting protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// LoginWithParam sets the LoginWithParam handler function
func (protocol *Protocol) LoginWithParam(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.loginWithParamHandler = handler
}

func (protocol *Protocol) handleLoginWithParam(packet nex.PacketInterface) {
	if protocol.loginWithParamHandler == nil {
		globals.Logger.Warning("TicketGranting::LoginWithParam not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	// Unsure what data is sent here, or how to trigger the console to send it
}
