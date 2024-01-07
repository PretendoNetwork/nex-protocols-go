// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func (protocol *Protocol) handleUpdateMatchmakeSessionPart(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdateMatchmakeSessionPart == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateMatchmakeSessionPart not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	updateMatchmakeSessionParam := match_making_types.NewUpdateMatchmakeSessionParam()
	err = updateMatchmakeSessionParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateMatchmakeSessionPart(fmt.Errorf("Failed to read updateMatchmakeSessionParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateMatchmakeSessionPart(nil, packet, callID, updateMatchmakeSessionParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
