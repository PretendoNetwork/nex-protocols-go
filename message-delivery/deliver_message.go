package message_delivery

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeliverMessage sets the DeliverMessage handler function
func (protocol *MessageDeliveryProtocol) DeliverMessage(handler func(err error, client *nex.Client, callID uint32, oUserMessage *nex.DataHolder)) {
	protocol.DeliverMessageHandler = handler
}

func (protocol *MessageDeliveryProtocol) HandleDeliverMessage(packet nex.PacketInterface) {
	if protocol.DeliverMessageHandler == nil {
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
		go protocol.DeliverMessageHandler(fmt.Errorf("Failed to read oUserMessage from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.DeliverMessageHandler(nil, client, callID, oUserMessage)
}
