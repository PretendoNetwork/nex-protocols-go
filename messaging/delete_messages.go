// Package messaging implements the Messaging protocol
package messaging

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

// DeleteMessages sets the DeleteMessages handler function
func (protocol *MessagingProtocol) DeleteMessages(handler func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient, lstMessagesToDelete []uint32)) {
	protocol.deleteMessagesHandler = handler
}

func (protocol *MessagingProtocol) handleDeleteMessages(packet nex.PacketInterface) {
	if protocol.deleteMessagesHandler == nil {
		globals.Logger.Warning("MessageDelivery::DeleteMessages not implemented")
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
		go protocol.deleteMessagesHandler(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	lstMessagesToDelete, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.deleteMessagesHandler(fmt.Errorf("Failed to read lstMessagesToDelete from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.deleteMessagesHandler(nil, client, callID, recipient.(*messaging_types.MessageRecipient), lstMessagesToDelete)
}
