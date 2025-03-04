// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUpdateMatchmakeSession(packet nex.PacketInterface) {
	if protocol.UpdateMatchmakeSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::UpdateMatchmakeSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var anyGathering match_making_types.GatheringHolder

	err := anyGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateMatchmakeSession(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, anyGathering)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateMatchmakeSession(nil, packet, callID, anyGathering)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
