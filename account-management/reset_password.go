// Package protocol implements the Account Management protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ResetPassword sets the ResetPassword handler function
func (protocol *Protocol) ResetPassword(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.resetPasswordHandler = handler
}

func (protocol *Protocol) handleResetPassword(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.resetPasswordHandler == nil {
		globals.Logger.Warning("AccountManagement::ResetPassword not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.resetPasswordHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
