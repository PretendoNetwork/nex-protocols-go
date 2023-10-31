// Package protocol implements the NAT Traversal protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRelaySignatureKey sets the GetRelaySignatureKey handler function
func (protocol *Protocol) GetRelaySignatureKey(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.getRelaySignatureKeyHandler = handler
}

func (protocol *Protocol) handleGetRelaySignatureKey(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getRelaySignatureKeyHandler == nil {
		globals.Logger.Warning("NATTraversal::GetRelaySignatureKey not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getRelaySignatureKeyHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
