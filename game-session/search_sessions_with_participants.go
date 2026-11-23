// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSearchSessionsWithParticipants(packet nex.PacketInterface) {
	if protocol.SearchSessionsWithParticipants == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "GameSession::SearchSessionsWithParticipants not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var gameSessionTypeID types.UInt32
	var participantIDs types.List[types.PID]

	var err error

	err = gameSessionTypeID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SearchSessionsWithParticipants(fmt.Errorf("failed to read gameSessionTypeID from parameters. %s", err.Error()), packet, callID, gameSessionTypeID, participantIDs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = participantIDs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SearchSessionsWithParticipants(fmt.Errorf("failed to read participantIDs from parameters. %s", err.Error()), packet, callID, gameSessionTypeID, participantIDs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchSessionsWithParticipants(nil, packet, callID, gameSessionTypeID, participantIDs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
