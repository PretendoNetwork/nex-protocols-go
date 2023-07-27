// Package protocol implements the NAT Traversal protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRelaySignatureKey sets the GetRelaySignatureKey handler function
func (protocol *Protocol) GetRelaySignatureKey(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getRelaySignatureKeyHandler = handler
}

func (protocol *Protocol) handleGetRelaySignatureKey(packet nex.PacketInterface) {
	if protocol.getRelaySignatureKeyHandler == nil {
		globals.Logger.Warning("NATTraversal::GetRelaySignatureKey not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getRelaySignatureKeyHandler(nil, client, callID)
}
