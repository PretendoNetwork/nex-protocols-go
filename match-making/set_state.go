// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	"github.com/PretendoNetwork/nex-protocols-go/v2/match-making/constants"
)

func (protocol *Protocol) handleSetState(packet nex.PacketInterface) {
	if protocol.SetState == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::SetState not implemented")

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
	var uiNewState constants.GatheringState

	var err error

	err = idGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetState(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, idGathering, uiNewState)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = uiNewState.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetState(fmt.Errorf("Failed to read uiNewState from parameters. %s", err.Error()), packet, callID, idGathering, uiNewState)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SetState(nil, packet, callID, idGathering, uiNewState)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
