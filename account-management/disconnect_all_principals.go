// Package protocol implements the Account Management protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DisconnectAllPrincipals sets the DisconnectAllPrincipals handler function
func (protocol *Protocol) DisconnectAllPrincipals(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.disconnectAllPrincipalsHandler = handler
}

func (protocol *Protocol) handleDisconnectAllPrincipals(packet nex.PacketInterface) {
	if protocol.disconnectAllPrincipalsHandler == nil {
		globals.Logger.Warning("AccountManagement::DisconnectAllPrincipals not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.disconnectAllPrincipalsHandler(nil, client, callID)
}
