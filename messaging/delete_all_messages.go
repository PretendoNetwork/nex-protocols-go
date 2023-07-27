// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

// DeleteAllMessages sets the DeleteAllMessages handler function
func (protocol *Protocol) DeleteAllMessages(handler func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient)) {
	protocol.deleteAllMessagesHandler = handler
}

func (protocol *Protocol) handleDeleteAllMessages(packet nex.PacketInterface) {
	if protocol.deleteAllMessagesHandler == nil {
		globals.Logger.Warning("Messaging::DeleteAllMessages not implemented")
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
		go protocol.deleteAllMessagesHandler(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.deleteAllMessagesHandler(nil, client, callID, recipient.(*messaging_types.MessageRecipient))
}
