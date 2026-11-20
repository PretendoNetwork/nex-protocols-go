// Package protocol implements the OfflineGameNotifications protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handlePollAnyOfflineNotifications(packet nex.PacketInterface) {
	if protocol.PollAnyOfflineNotifications == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "OfflineGameNotifications::PollAnyOfflineNotifications not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.PollAnyOfflineNotifications(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
