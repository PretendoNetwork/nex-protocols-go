// Package protocol implements the Utility protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetAssociatedNexUniqueIDWithMyPrincipalID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetAssociatedNexUniqueIDWithMyPrincipalID == nil {
		globals.Logger.Warning("Utility::GetAssociatedNexUniqueIDWithMyPrincipalID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.GetAssociatedNexUniqueIDWithMyPrincipalID(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
