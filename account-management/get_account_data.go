// Package protocol implements the Account Management protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAccountData sets the GetAccountData handler function
func (protocol *Protocol) GetAccountData(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.getAccountDataHandler = handler
}

func (protocol *Protocol) handleGetAccountData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getAccountDataHandler == nil {
		globals.Logger.Warning("AccountManagement::GetAccountData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.getAccountDataHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
