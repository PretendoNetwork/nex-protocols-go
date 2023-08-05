// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// WithdrawMatchmaking sets the WithdrawMatchmaking handler function
func (protocol *Protocol) WithdrawMatchmaking(handler func(err error, client *nex.Client, callID uint32, requestID uint64) uint32) {
	protocol.withdrawMatchmakingHandler = handler
}

func (protocol *Protocol) handleWithdrawMatchmaking(packet nex.PacketInterface) {
	if protocol.withdrawMatchmakingHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::WithdrawMatchmaking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	requestID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.updateProgressScoreHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.withdrawMatchmakingHandler(nil, client, callID, requestID)
}
