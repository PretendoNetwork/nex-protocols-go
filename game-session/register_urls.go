// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleRegisterURLs(packet nex.PacketInterface) {
	if protocol.RegisterURLs == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "GameSession::RegisterURLs not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var stationURLs types.List[types.StationURL]

	err := stationURLs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RegisterURLs(fmt.Errorf("failed to read stationURLs from parameters. %s", err.Error()), packet, callID, stationURLs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RegisterURLs(nil, packet, callID, stationURLs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
