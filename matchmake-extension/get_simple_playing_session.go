// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetSimplePlayingSession(packet nex.PacketInterface) {
	if protocol.GetSimplePlayingSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::GetSimplePlayingSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var listPID types.List[types.PID]
	var includeLoginUser types.Bool

	var err error

	err = listPID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetSimplePlayingSession(fmt.Errorf("Failed to read listPID from parameters. %s", err.Error()), packet, callID, listPID, includeLoginUser)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = includeLoginUser.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetSimplePlayingSession(fmt.Errorf("Failed to read includeLoginUser from parameters. %s", err.Error()), packet, callID, listPID, includeLoginUser)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetSimplePlayingSession(nil, packet, callID, listPID, includeLoginUser)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
