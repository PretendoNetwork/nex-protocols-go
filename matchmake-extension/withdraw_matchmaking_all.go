// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// WithdrawMatchmakingAll sets the WithdrawMatchmakingAll handler function
func (protocol *Protocol) WithdrawMatchmakingAll(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.withdrawMatchmakingAllHandler = handler
}

func (protocol *Protocol) handleWithdrawMatchmakingAll(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.withdrawMatchmakingAllHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::WithdrawMatchmakingAll not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.withdrawMatchmakingAllHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
