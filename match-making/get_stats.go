// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetStats(packet nex.PacketInterface) {
	if protocol.GetStats == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::GetStats not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var idGathering types.UInt32
	var lstParticipants types.List[types.PID]
	var lstColumns types.Buffer

	var err error

	err = idGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetStats(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, idGathering, lstParticipants, lstColumns)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = lstParticipants.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetStats(fmt.Errorf("Failed to read lstParticipants from parameters. %s", err.Error()), packet, callID, idGathering, lstParticipants, lstColumns)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = lstColumns.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetStats(fmt.Errorf("Failed to read lstColumns from parameters. %s", err.Error()), packet, callID, idGathering, lstParticipants, lstColumns)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetStats(nil, packet, callID, idGathering, lstParticipants, lstColumns)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
