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
	var errorCode uint32

	if protocol.DeliverMessage == nil {
		globals.Logger.Warning("MessageDelivery::DeliverMessage not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	oUserMessage := types.NewAnyDataHolder()
	err = oUserMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DeliverMessage(fmt.Errorf("Failed to read oUserMessage from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DeliverMessage(nil, packet, callID, oUserMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
