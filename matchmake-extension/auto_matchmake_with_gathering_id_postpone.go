// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleAutoMatchmakeWithGatheringIDPostpone(packet nex.PacketInterface) {
	if protocol.AutoMatchmakeWithGatheringIDPostpone == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::AutoMatchmakeWithGatheringIDPostpone not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var lstGID types.List[types.UInt32]
	var anyGathering match_making_types.GatheringHolder
	var strMessage types.String

	var err error

	err = lstGID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AutoMatchmakeWithGatheringIDPostpone(fmt.Errorf("Failed to read lstGID from parameters. %s", err.Error()), packet, callID, lstGID, anyGathering, strMessage)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = anyGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AutoMatchmakeWithGatheringIDPostpone(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, lstGID, anyGathering, strMessage)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AutoMatchmakeWithGatheringIDPostpone(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, lstGID, anyGathering, strMessage)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AutoMatchmakeWithGatheringIDPostpone(nil, packet, callID, lstGID, anyGathering, strMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
