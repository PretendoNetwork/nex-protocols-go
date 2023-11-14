// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func (protocol *Protocol) handleFindMatchmakeSessionByParticipant(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.FindMatchmakeSessionByParticipant == nil {
		globals.Logger.Warning("MatchmakeExtension::FindMatchmakeSessionByParticipant not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(match_making_types.NewFindMatchmakeSessionByParticipantParam())
	if err != nil {
		errorCode = protocol.FindMatchmakeSessionByParticipant(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.FindMatchmakeSessionByParticipant(nil, packet, callID, param.(*match_making_types.FindMatchmakeSessionByParticipantParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
