// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

func (protocol *Protocol) handleRetrieveMessages(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.RetrieveMessages == nil {
		globals.Logger.Warning("Messaging::RetrieveMessages not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	recipient := messaging_types.NewMessageRecipient()
	err = recipient.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RetrieveMessages(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstMsgIDs := types.NewList[*types.PrimitiveU32]()
	lstMsgIDs.Type = types.NewPrimitiveU32(0)
	err = lstMsgIDs.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RetrieveMessages(fmt.Errorf("Failed to read lstMsgIDs from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bLeaveOnServer := types.NewPrimitiveBool(false)
	err = bLeaveOnServer.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RetrieveMessages(fmt.Errorf("Failed to read bLeaveOnServer from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
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
