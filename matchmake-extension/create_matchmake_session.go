// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleCreateMatchmakeSession(packet nex.PacketInterface) {
	if protocol.CreateMatchmakeSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::CreateMatchmakeSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	endpoint := packet.Sender().Endpoint()
	matchmakingVersion := endpoint.LibraryVersions().MatchMaking

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var anyGathering match_making_types.GatheringHolder
	var strMessage types.String
	var participationCount types.UInt16

	var err error

	err = anyGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateMatchmakeSession(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, anyGathering, strMessage, participationCount)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateMatchmakeSession(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, anyGathering, strMessage, participationCount)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	if matchmakingVersion.GreaterOrEqual("3.4.0") {
		err = participationCount.ExtractFrom(parametersStream)
		if err != nil {
			_, rmcError := protocol.CreateMatchmakeSession(fmt.Errorf("Failed to read participationCount from parameters. %s", err.Error()), packet, callID, anyGathering, strMessage, participationCount)
			if rmcError != nil {
				globals.RespondError(packet, ProtocolID, rmcError)
			}

			return
		}
	}

	rmcMessage, rmcError := protocol.CreateMatchmakeSession(nil, packet, callID, anyGathering, strMessage, participationCount)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
