// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

// RetrieveMessages sets the RetrieveMessages handler function
func (protocol *Protocol) RetrieveMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMsgIDs []uint32, bLeaveOnServer bool) uint32) {
	protocol.retrieveMessagesHandler = handler
}

func (protocol *Protocol) handleRetrieveMessages(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.retrieveMessagesHandler == nil {
		globals.Logger.Warning("Messaging::RetrieveMessages not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	recipient, err := parametersStream.ReadStructure(messaging_types.NewMessageRecipient())
	if err != nil {
		errorCode = protocol.retrieveMessagesHandler(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstMsgIDs, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.retrieveMessagesHandler(fmt.Errorf("Failed to read lstMsgIDs from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bLeaveOnServer, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.retrieveMessagesHandler(fmt.Errorf("Failed to read bLeaveOnServer from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.retrieveMessagesHandler(nil, packet, callID, recipient.(*messaging_types.MessageRecipient), lstMsgIDs, bLeaveOnServer)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
