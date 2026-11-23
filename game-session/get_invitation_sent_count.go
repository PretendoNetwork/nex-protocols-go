// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetInvitationSentCount(packet nex.PacketInterface) {
	if protocol.GetInvitationSentCount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "GameSession::GetInvitationSentCount not implemented")

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

	err := gameSessionTypeID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetInvitationSentCount(fmt.Errorf("failed to read gameSessionTypeID from parameters. %s", err.Error()), packet, callID, gameSessionTypeID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetInvitationSentCount(nil, packet, callID, gameSessionTypeID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
