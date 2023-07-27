// Package protocol implements the Message Deliver protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeliverMessage sets the DeliverMessage handler function
func (protocol *Protocol) DeliverMessage(handler func(err error, client *nex.Client, callID uint32, oUserMessage *nex.DataHolder)) {
	protocol.deliverMessageHandler = handler
}

func (protocol *Protocol) handleDeliverMessage(packet nex.PacketInterface) {
	if protocol.deliverMessageHandler == nil {
		globals.Logger.Warning("MessageDelivery::DeliverMessage not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	oUserMessage, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.deliverMessageHandler(fmt.Errorf("Failed to read oUserMessage from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.deliverMessageHandler(nil, client, callID, oUserMessage)
}
