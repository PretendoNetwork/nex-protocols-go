// Package protocol implements the Ticket Granting protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleLoginWithContext(packet nex.PacketInterface) {
	if protocol.LoginWithContext == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "TicketGranting::LoginWithContext not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	// * Unsure what data is sent here, or how to trigger the console to send it
}
