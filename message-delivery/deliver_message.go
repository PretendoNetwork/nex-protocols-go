// Package protocol implements the Message Deliver protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeliverMessage(packet nex.PacketInterface) {
	var err error

	if protocol.DeliverMessage == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MessageDelivery::DeliverMessage not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	oUserMessage := types.NewAnyDataHolder()
	err = oUserMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeliverMessage(fmt.Errorf("Failed to read oUserMessage from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeliverMessage(nil, packet, callID, oUserMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
