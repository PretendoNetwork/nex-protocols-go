// Package protocol implements the OfflineGameNotifications protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handlePollSpecificOfflineNotifications(packet nex.PacketInterface) {
	if protocol.PollSpecificOfflineNotifications == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "OfflineGameNotifications::PollSpecificOfflineNotifications not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var majorTypes types.List[types.UInt32]

	err := majorTypes.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PollSpecificOfflineNotifications(fmt.Errorf("failed to read majortypes from parameters. %s", err.Error()), packet, callID, majorTypes)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PollSpecificOfflineNotifications(nil, packet, callID, majorTypes)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
