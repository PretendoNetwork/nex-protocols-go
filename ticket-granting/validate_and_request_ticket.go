// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleValidateAndRequestTicket(packet nex.PacketInterface) {
	if protocol.ValidateAndRequestTicket == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "TicketGranting::ValidateAndRequestTicket not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var strUserName types.String

	err := strUserName.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ValidateAndRequestTicket(fmt.Errorf("Failed to read strUserName from parameters. %s", err.Error()), packet, callID, strUserName)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ValidateAndRequestTicket(nil, packet, callID, strUserName)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
