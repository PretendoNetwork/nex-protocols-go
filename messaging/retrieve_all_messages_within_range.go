// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

// RetrieveAllMessagesWithinRange sets the RetrieveAllMessagesWithinRange handler function
func (protocol *Protocol) RetrieveAllMessagesWithinRange(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange) uint32) {
	protocol.retrieveAllMessagesWithinRangeHandler = handler
}

func (protocol *Protocol) handleRetrieveAllMessagesWithinRange(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.retrieveAllMessagesWithinRangeHandler == nil {
		globals.Logger.Warning("Messaging::RetrieveAllMessagesWithinRange not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	recipient, err := parametersStream.ReadStructure(messaging_types.NewMessageRecipient())
	if err != nil {
		errorCode = protocol.retrieveAllMessagesWithinRangeHandler(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.retrieveAllMessagesWithinRangeHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.retrieveAllMessagesWithinRangeHandler(nil, packet, callID, recipient.(*messaging_types.MessageRecipient), resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
