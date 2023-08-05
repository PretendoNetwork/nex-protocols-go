// Package protocol implements the Utility protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAssociatedNexUniqueIDsWithMyPrincipalID sets the GetAssociatedNexUniqueIDsWithMyPrincipalID handler function
func (protocol *Protocol) GetAssociatedNexUniqueIDsWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.getAssociatedNexUniqueIDsWithMyPrincipalIDHandler = handler
}

func (protocol *Protocol) handleGetAssociatedNexUniqueIDsWithMyPrincipalID(packet nex.PacketInterface) {
	if protocol.getAssociatedNexUniqueIDsWithMyPrincipalIDHandler == nil {
		globals.Logger.Warning("Utility::GetAssociatedNexUniqueIDsWithMyPrincipalID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getAssociatedNexUniqueIDsWithMyPrincipalIDHandler(nil, client, callID)
}
