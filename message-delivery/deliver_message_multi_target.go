// Package protocol implements the Message Delivery protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleDeliverMessageMultiTarget(packet nex.PacketInterface) {
	if protocol.DeliverMessageMultiTarget == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MessageDelivery::DeliverMessageMultiTarget not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var lstTarget types.List[types.PID]
	var oUserMessage types.DataHolder

	var err error

	err = lstTarget.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeliverMessageMultiTarget(fmt.Errorf("Failed to read lstTarget from parameters. %s", err.Error()), packet, callID, lstTarget, oUserMessage)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = oUserMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeliverMessageMultiTarget(fmt.Errorf("Failed to read oUserMessage from parameters. %s", err.Error()), packet, callID, lstTarget, oUserMessage)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeliverMessageMultiTarget(nil, packet, callID, lstTarget, oUserMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
