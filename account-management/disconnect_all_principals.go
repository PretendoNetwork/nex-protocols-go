// Package protocol implements the Account Management protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DisconnectAllPrincipals sets the DisconnectAllPrincipals handler function
func (protocol *Protocol) DisconnectAllPrincipals(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.disconnectAllPrincipalsHandler = handler
}

func (protocol *Protocol) handleDisconnectAllPrincipals(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.disconnectAllPrincipalsHandler == nil {
		globals.Logger.Warning("AccountManagement::DisconnectAllPrincipals not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.disconnectAllPrincipalsHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
