// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// WithdrawMatchmaking sets the WithdrawMatchmaking handler function
func (protocol *Protocol) WithdrawMatchmaking(handler func(err error, packet nex.PacketInterface, callID uint32, requestID uint64) uint32) {
	protocol.withdrawMatchmakingHandler = handler
}

func (protocol *Protocol) handleWithdrawMatchmaking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.withdrawMatchmakingHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::WithdrawMatchmaking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	requestID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.updateProgressScoreHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.withdrawMatchmakingHandler(nil, packet, callID, requestID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
