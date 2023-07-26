// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// WithdrawMatchmakingAll sets the WithdrawMatchmakingAll handler function
func (protocol *Protocol) WithdrawMatchmakingAll(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.withdrawMatchmakingAllHandler = handler
}

func (protocol *Protocol) handleWithdrawMatchmakingAll(packet nex.PacketInterface) {
	if protocol.withdrawMatchmakingAllHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::WithdrawMatchmakingAll not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.withdrawMatchmakingAllHandler(nil, client, callID)
}
