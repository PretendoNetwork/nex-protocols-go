// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	"github.com/PretendoNetwork/nex-protocols-go/v2/ticket-granting/types"
)

func (protocol *Protocol) handleValidateAndRequestTicketWithParam(packet nex.PacketInterface) {
	if protocol.ValidateAndRequestTicketWithParam == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "TicketGranting::ValidateAndRequestTicketWithParam not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var param types.ValidateAndRequestTicketParam

	err := param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ValidateAndRequestTicketWithParam(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, param)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ValidateAndRequestTicketWithParam(nil, packet, callID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
