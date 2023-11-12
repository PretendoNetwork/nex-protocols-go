// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

// GetNumberOfMessages sets the GetNumberOfMessages handler function
func (protocol *Protocol) GetNumberOfMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) uint32) {
	protocol.getNumberOfMessagesHandler = handler
}

func (protocol *Protocol) handleGetNumberOfMessages(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getNumberOfMessagesHandler == nil {
		globals.Logger.Warning("Messaging::GetNumberOfMessages not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	recipient, err := parametersStream.ReadStructure(messaging_types.NewMessageRecipient())
	if err != nil {
		errorCode = protocol.getNumberOfMessagesHandler(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getNumberOfMessagesHandler(nil, packet, callID, recipient.(*messaging_types.MessageRecipient))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
