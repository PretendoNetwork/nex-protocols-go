// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// AutoMatchmakeWithParamPostpone sets the AutoMatchmakeWithParamPostpone handler function
func (protocol *Protocol) AutoMatchmakeWithParamPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) uint32) {
	protocol.autoMatchmakeWithParamPostponeHandler = handler
}

func (protocol *Protocol) handleAutoMatchmakeWithParamPostpone(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.autoMatchmakeWithParamPostponeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakeWithParamPostpone not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	autoMatchmakeParam, err := parametersStream.ReadStructure(match_making_types.NewAutoMatchmakeParam())
	if err != nil {
		errorCode = protocol.autoMatchmakeWithParamPostponeHandler(fmt.Errorf("Failed to read autoMatchmakeParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.autoMatchmakeWithParamPostponeHandler(nil, packet, callID, autoMatchmakeParam.(*match_making_types.AutoMatchmakeParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
