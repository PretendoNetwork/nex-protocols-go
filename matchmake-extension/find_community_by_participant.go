// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleFindCommunityByParticipant(packet nex.PacketInterface) {
	if protocol.FindCommunityByParticipant == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::FindCommunityByParticipant not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var pid types.PID
	var resultRange types.ResultRange

	var err error

	err = pid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindCommunityByParticipant(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, pid, resultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindCommunityByParticipant(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, pid, resultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.FindCommunityByParticipant(nil, packet, callID, pid, resultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
