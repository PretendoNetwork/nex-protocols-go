// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleWithdrawMatchmaking(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.WithdrawMatchmaking == nil {
		globals.Logger.Warning("MatchmakeExtension::WithdrawMatchmaking not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	requestID := types.NewPrimitiveU64(0)
	err = requestID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateProgressScore(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.WithdrawMatchmaking(nil, packet, callID, requestID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
