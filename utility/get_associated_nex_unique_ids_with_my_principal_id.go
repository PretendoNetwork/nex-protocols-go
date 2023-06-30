// Package utility implements the Utility NEX protocol
package utility

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAssociatedNexUniqueIDsWithMyPrincipalID sets the GetAssociatedNexUniqueIDsWithMyPrincipalID handler function
func (protocol *UtilityProtocol) GetAssociatedNexUniqueIDsWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetAssociatedNexUniqueIDsWithMyPrincipalIDHandler = handler
}

func (protocol *UtilityProtocol) handleGetAssociatedNexUniqueIDsWithMyPrincipalID(packet nex.PacketInterface) {
	if protocol.GetAssociatedNexUniqueIDsWithMyPrincipalIDHandler == nil {
		globals.Logger.Warning("Utility::GetAssociatedNexUniqueIDsWithMyPrincipalID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.GetAssociatedNexUniqueIDsWithMyPrincipalIDHandler(nil, client, callID)
}
