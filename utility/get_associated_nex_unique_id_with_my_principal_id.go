// Package protocol implements the Utility protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAssociatedNexUniqueIDWithMyPrincipalID sets the GetAssociatedNexUniqueIDWithMyPrincipalID handler function
func (protocol *Protocol) GetAssociatedNexUniqueIDWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getAssociatedNexUniqueIDWithMyPrincipalIDHandler = handler
}

func (protocol *Protocol) handleGetAssociatedNexUniqueIDWithMyPrincipalID(packet nex.PacketInterface) {
	if protocol.getAssociatedNexUniqueIDWithMyPrincipalIDHandler == nil {
		globals.Logger.Warning("Utility::GetAssociatedNexUniqueIDWithMyPrincipalID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getAssociatedNexUniqueIDWithMyPrincipalIDHandler(nil, client, callID)
}
