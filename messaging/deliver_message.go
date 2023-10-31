// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeliverMessage sets the DeliverMessage handler function
func (protocol *Protocol) DeliverMessage(handler func(err error, packet nex.PacketInterface, callID uint32, oUserMessage *nex.DataHolder) uint32) {
	protocol.deliverMessageHandler = handler
}

func (protocol *Protocol) handleDeliverMessage(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deliverMessageHandler == nil {
		globals.Logger.Warning("Messaging::DeliverMessage not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	oUserMessage, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.deleteMessagesHandler(fmt.Errorf("Failed to read oUserMessage from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.deliverMessageHandler(nil, packet, callID, oUserMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
