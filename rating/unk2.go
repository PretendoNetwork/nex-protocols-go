// Package protocol implements the Rating protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	rating_types "github.com/PretendoNetwork/nex-protocols-go/v2/rating/types"
)

// TODO - Find name if possible
func (protocol *Protocol) handleUnk2(packet nex.PacketInterface) {
	if protocol.Unk2 == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Rating::Unk2 not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var sessionToken rating_types.RatingSessionToken

	var err error

	err = sessionToken.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.Unk2(fmt.Errorf("Failed to read sessionToken from parameters. %s", err.Error()), packet, callID, sessionToken)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.Unk2(nil, packet, callID, sessionToken)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
