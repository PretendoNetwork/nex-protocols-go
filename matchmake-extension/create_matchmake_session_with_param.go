// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func (protocol *Protocol) handleCreateMatchmakeSessionWithParam(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.CreateMatchmakeSessionWithParam == nil {
		globals.Logger.Warning("MatchmakeExtension::CreateMatchmakeSessionWithParam not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	createMatchmakeSessionParam := match_making_types.NewCreateMatchmakeSessionParam()
	err = createMatchmakeSessionParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.CreateMatchmakeSessionWithParam(fmt.Errorf("Failed to read createMatchmakeSessionParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.CreateMatchmakeSessionWithParam(nil, packet, callID, createMatchmakeSessionParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
