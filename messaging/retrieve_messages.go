// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

func (protocol *Protocol) handleRetrieveMessages(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.RetrieveMessages == nil {
		globals.Logger.Warning("Messaging::RetrieveMessages not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	recipient, err := nex.StreamReadStructure(parametersStream, messaging_types.NewMessageRecipient())
	if err != nil {
		_, errorCode = protocol.RetrieveMessages(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstMsgIDs, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		_, errorCode = protocol.RetrieveMessages(fmt.Errorf("Failed to read lstMsgIDs from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bLeaveOnServer, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.RetrieveMessages(fmt.Errorf("Failed to read bLeaveOnServer from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RetrieveMessages(nil, packet, callID, recipient, lstMsgIDs, bLeaveOnServer)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
