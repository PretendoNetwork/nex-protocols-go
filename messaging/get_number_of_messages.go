// Package messaging implements the Messaging protocol
package messaging

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

// GetNumberOfMessages sets the GetNumberOfMessages handler function
func (protocol *MessagingProtocol) GetNumberOfMessages(handler func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient)) {
	protocol.getNumberOfMessagesHandler = handler
}

func (protocol *MessagingProtocol) handleGetNumberOfMessages(packet nex.PacketInterface) {
	if protocol.getNumberOfMessagesHandler == nil {
		globals.Logger.Warning("MessageDelivery::GetNumberOfMessages not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	recipient, err := parametersStream.ReadStructure(messaging_types.NewMessageRecipient())
	if err != nil {
		go protocol.getNumberOfMessagesHandler(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getNumberOfMessagesHandler(nil, client, callID, recipient.(*messaging_types.MessageRecipient))
}
