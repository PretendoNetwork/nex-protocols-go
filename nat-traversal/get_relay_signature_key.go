// Package nat_traversal implements the NAT Traversal NEX protocol
package nat_traversal

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRelaySignatureKey sets the GetRelaySignatureKey handler function
func (protocol *NATTraversalProtocol) GetRelaySignatureKey(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetRelaySignatureKeyHandler = handler
}

func (protocol *NATTraversalProtocol) handleGetRelaySignatureKey(packet nex.PacketInterface) {
	if protocol.GetRelaySignatureKeyHandler == nil {
		globals.Logger.Warning("NATTraversal::GetRelaySignatureKey not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.GetRelaySignatureKeyHandler(nil, client, callID)
}
