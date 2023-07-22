// Package account_management implements the Account Management NEX protocol
package account_management

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RetrieveAccount sets the RetrieveAccount handler function
func (protocol *AccountManagementProtocol) RetrieveAccount(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.retrieveAccountHandler = handler
}

func (protocol *AccountManagementProtocol) handleRetrieveAccount(packet nex.PacketInterface) {
	if protocol.retrieveAccountHandler == nil {
		globals.Logger.Warning("AccountManagement::RetrieveAccount not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.retrieveAccountHandler(nil, client, callID)
}
