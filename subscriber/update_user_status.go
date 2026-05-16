// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/v2/subscriber/types"
)

func (protocol *Protocol) handleUpdateUserStatus(packet nex.PacketInterface) {
	if protocol.UpdateUserStatus == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Subscriber::UpdateUserStatus not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var param types.List[subscriber_types.SubscriberUserStatusParam]
	var isNotify types.List[types.UInt8]

	var err error

	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateUserStatus(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, param, isNotify)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = isNotify.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateUserStatus(fmt.Errorf("Failed to read isNotify from parameters. %s", err.Error()), packet, callID, param, isNotify)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateUserStatus(nil, packet, callID, param, isNotify)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
