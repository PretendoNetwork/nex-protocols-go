// Package protocol implements the Utility protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAssociatedNexUniqueIDWithMyPrincipalID sets the GetAssociatedNexUniqueIDWithMyPrincipalID handler function
func (protocol *Protocol) GetAssociatedNexUniqueIDWithMyPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.getAssociatedNexUniqueIDWithMyPrincipalIDHandler = handler
}

func (protocol *Protocol) handleGetAssociatedNexUniqueIDWithMyPrincipalID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getAssociatedNexUniqueIDWithMyPrincipalIDHandler == nil {
		globals.Logger.Warning("Utility::GetAssociatedNexUniqueIDWithMyPrincipalID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getAssociatedNexUniqueIDWithMyPrincipalIDHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
