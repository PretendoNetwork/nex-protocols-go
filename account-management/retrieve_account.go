// Package protocol implements the Account Management protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RetrieveAccount sets the RetrieveAccount handler function
func (protocol *Protocol) RetrieveAccount(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.retrieveAccountHandler = handler
}

func (protocol *Protocol) handleRetrieveAccount(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.retrieveAccountHandler == nil {
		globals.Logger.Warning("AccountManagement::RetrieveAccount not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.retrieveAccountHandler(nil, client, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
