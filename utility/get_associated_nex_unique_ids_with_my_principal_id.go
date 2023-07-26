// Package utility implements the Utility NEX protocol
package utility

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAssociatedNexUniqueIDsWithMyPrincipalID sets the GetAssociatedNexUniqueIDsWithMyPrincipalID handler function
func (protocol *UtilityProtocol) GetAssociatedNexUniqueIDsWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getAssociatedNexUniqueIDsWithMyPrincipalIDHandler = handler
}

func (protocol *UtilityProtocol) handleGetAssociatedNexUniqueIDsWithMyPrincipalID(packet nex.PacketInterface) {
	if protocol.getAssociatedNexUniqueIDsWithMyPrincipalIDHandler == nil {
		globals.Logger.Warning("Utility::GetAssociatedNexUniqueIDsWithMyPrincipalID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getAssociatedNexUniqueIDsWithMyPrincipalIDHandler(nil, client, callID)
}
