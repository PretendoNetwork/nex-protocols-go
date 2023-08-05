// Package protocol implements the Ticket Granting protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// LoginWithContext sets the LoginWithContext handler function
func (protocol *Protocol) LoginWithContext(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.loginWithContextHandler = handler
}

func (protocol *Protocol) handleLoginWithContext(packet nex.PacketInterface) {
	if protocol.loginWithContextHandler == nil {
		globals.Logger.Warning("TicketGranting::LoginWithContext not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	// Unsure what data is sent here, or how to trigger the console to send it
}
