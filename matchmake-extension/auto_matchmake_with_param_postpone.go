// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func (protocol *Protocol) handleAutoMatchmakeWithParamPostpone(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AutoMatchmakeWithParamPostpone == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmakeWithParamPostpone not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	autoMatchmakeParam, err := nex.StreamReadStructure(parametersStream, match_making_types.NewAutoMatchmakeParam())
	if err != nil {
		_, errorCode = protocol.AutoMatchmakeWithParamPostpone(fmt.Errorf("Failed to read autoMatchmakeParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AutoMatchmakeWithParamPostpone(nil, packet, callID, autoMatchmakeParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
