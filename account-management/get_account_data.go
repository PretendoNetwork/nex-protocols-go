// Package protocol implements the Account Management protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAccountData sets the GetAccountData handler function
func (protocol *Protocol) GetAccountData(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.getAccountDataHandler = handler
}

func (protocol *Protocol) handleGetAccountData(packet nex.PacketInterface) {
	if protocol.getAccountDataHandler == nil {
		globals.Logger.Warning("AccountManagement::GetAccountData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getAccountDataHandler(nil, client, callID)
}
