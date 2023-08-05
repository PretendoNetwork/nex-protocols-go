// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

// GetMessagesHeaders sets the GetMessagesHeaders handler function
func (protocol *Protocol) GetMessagesHeaders(handler func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange) uint32) {
	protocol.getMessagesHeadersHandler = handler
}

func (protocol *Protocol) handleGetMessagesHeaders(packet nex.PacketInterface) {
	if protocol.getMessagesHeadersHandler == nil {
		globals.Logger.Warning("Messaging::GetMessagesHeaders not implemented")
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
		go protocol.getMessagesHeadersHandler(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.getMessagesHeadersHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.getMessagesHeadersHandler(nil, client, callID, recipient.(*messaging_types.MessageRecipient), resultRange.(*nex.ResultRange))
}
