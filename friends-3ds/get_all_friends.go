// Package protocol implements the Friends 3DS protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetAllFriends(packet nex.PacketInterface) {
	if protocol.GetAllFriends == nil {
		globals.Logger.Warning("Friends3DS::GetAllFriends not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.GetAllFriends(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
