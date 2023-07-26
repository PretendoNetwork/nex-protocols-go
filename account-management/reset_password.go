// Package protocol implements the Account Management protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ResetPassword sets the ResetPassword handler function
func (protocol *Protocol) ResetPassword(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.resetPasswordHandler = handler
}

func (protocol *Protocol) handleResetPassword(packet nex.PacketInterface) {
	if protocol.resetPasswordHandler == nil {
		globals.Logger.Warning("AccountManagement::ResetPassword not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.resetPasswordHandler(nil, client, callID)
}
