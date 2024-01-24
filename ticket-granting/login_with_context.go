// Package protocol implements the Ticket Granting protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleLoginWithContext(packet nex.PacketInterface) {
	if protocol.LoginWithContext == nil {
		globals.Logger.Warning("TicketGranting::LoginWithContext not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	// * Unsure what data is sent here, or how to trigger the console to send it
}
