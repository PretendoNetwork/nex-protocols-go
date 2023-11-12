// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// RequestMatchmaking sets the RequestMatchmaking handler function
func (protocol *Protocol) RequestMatchmaking(handler func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) uint32) {
	protocol.requestMatchmakingHandler = handler
}

func (protocol *Protocol) handleRequestMatchmaking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.requestMatchmakingHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::RequestMatchmaking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	autoMatchmakeParam, err := parametersStream.ReadStructure(match_making_types.NewAutoMatchmakeParam())
	if err != nil {
		errorCode = protocol.requestMatchmakingHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.requestMatchmakingHandler(nil, packet, callID, autoMatchmakeParam.(*match_making_types.AutoMatchmakeParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
