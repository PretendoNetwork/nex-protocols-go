// Package utility implements the Utility NEX protocol
package utility

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAssociatedNexUniqueIDWithMyPrincipalID sets the GetAssociatedNexUniqueIDWithMyPrincipalID handler function
func (protocol *UtilityProtocol) GetAssociatedNexUniqueIDWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetAssociatedNexUniqueIDWithMyPrincipalIDHandler = handler
}

func (protocol *UtilityProtocol) handleGetAssociatedNexUniqueIDWithMyPrincipalID(packet nex.PacketInterface) {
	if protocol.GetAssociatedNexUniqueIDWithMyPrincipalIDHandler == nil {
		globals.Logger.Warning("Utility::GetAssociatedNexUniqueIDWithMyPrincipalID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.GetAssociatedNexUniqueIDWithMyPrincipalIDHandler(nil, client, callID)
}
