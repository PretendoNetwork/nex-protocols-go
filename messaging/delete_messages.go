// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

// DeleteMessages sets the DeleteMessages handler function
func (protocol *Protocol) DeleteMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMessagesToDelete []uint32) uint32) {
	protocol.deleteMessagesHandler = handler
}

func (protocol *Protocol) handleDeleteMessages(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deleteMessagesHandler == nil {
		globals.Logger.Warning("Messaging::DeleteMessages not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	recipient, err := parametersStream.ReadStructure(messaging_types.NewMessageRecipient())
	if err != nil {
		errorCode = protocol.deleteMessagesHandler(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstMessagesToDelete, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.deleteMessagesHandler(fmt.Errorf("Failed to read lstMessagesToDelete from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.deleteMessagesHandler(nil, packet, callID, recipient.(*messaging_types.MessageRecipient), lstMessagesToDelete)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
