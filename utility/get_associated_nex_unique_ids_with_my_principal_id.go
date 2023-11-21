// Package protocol implements the Utility protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetAssociatedNexUniqueIDsWithMyPrincipalID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetAssociatedNexUniqueIDsWithMyPrincipalID == nil {
		globals.Logger.Warning("Utility::GetAssociatedNexUniqueIDsWithMyPrincipalID not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.GetAssociatedNexUniqueIDsWithMyPrincipalID(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
