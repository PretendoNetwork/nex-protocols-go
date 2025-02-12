// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleFindByDescriptionLike(packet nex.PacketInterface) {
	if protocol.FindByDescriptionLike == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::FindByDescriptionLike not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var strDescriptionLike types.String
	var resultRange types.ResultRange

	var err error

	err = strDescriptionLike.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByDescriptionLike(fmt.Errorf("Failed to read strDescriptionLike from parameters. %s", err.Error()), packet, callID, strDescriptionLike, resultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByDescriptionLike(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, strDescriptionLike, resultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.FindByDescriptionLike(nil, packet, callID, strDescriptionLike, resultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
