// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func (protocol *Protocol) handleUpdateMatchmakeSessionPart(packet nex.PacketInterface) {
	var err error

	if protocol.UpdateMatchmakeSessionPart == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::UpdateMatchmakeSessionPart not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	updateMatchmakeSessionParam := match_making_types.NewUpdateMatchmakeSessionParam()
	err = updateMatchmakeSessionParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateMatchmakeSessionPart(fmt.Errorf("Failed to read updateMatchmakeSessionParam from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateMatchmakeSessionPart(nil, packet, callID, updateMatchmakeSessionParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
