// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

// GetNumberOfMessages sets the GetNumberOfMessages handler function
func (protocol *Protocol) GetNumberOfMessages(handler func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient) uint32) {
	protocol.getNumberOfMessagesHandler = handler
}

func (protocol *Protocol) handleGetNumberOfMessages(packet nex.PacketInterface) {
	if protocol.getNumberOfMessagesHandler == nil {
		globals.Logger.Warning("Messaging::GetNumberOfMessages not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
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
