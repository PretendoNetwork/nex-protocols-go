// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// WithdrawMatchmakingAll sets the WithdrawMatchmakingAll handler function
func (protocol *MatchmakeExtensionProtocol) WithdrawMatchmakingAll(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.withdrawMatchmakingAllHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleWithdrawMatchmakingAll(packet nex.PacketInterface) {
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
