// Package account_management implements the Account Management NEX protocol
package account_management

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAccountData sets the GetAccountData handler function
func (protocol *AccountManagementProtocol) GetAccountData(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getAccountDataHandler = handler
}

func (protocol *AccountManagementProtocol) handleGetAccountData(packet nex.PacketInterface) {
	if protocol.getAccountDataHandler == nil {
		globals.Logger.Warning("AccountManagement::GetAccountData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getAccountDataHandler(nil, client, callID)
}
