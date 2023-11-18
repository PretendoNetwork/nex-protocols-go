// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

func (protocol *Protocol) handleDeleteMessages(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.DeleteMessages == nil {
		globals.Logger.Warning("Messaging::DeleteMessages not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	recipient, err := nex.StreamReadStructure(parametersStream, messaging_types.NewMessageRecipient())
	if err != nil {
		_, errorCode = protocol.DeleteMessages(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstMessagesToDelete, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		_, errorCode = protocol.DeleteMessages(fmt.Errorf("Failed to read lstMessagesToDelete from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DeleteMessages(nil, packet, callID, recipient, lstMessagesToDelete)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
