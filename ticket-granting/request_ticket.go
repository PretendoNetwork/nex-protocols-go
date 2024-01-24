// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRequestTicket(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.RequestTicket == nil {
		globals.Logger.Warning("TicketGranting::RequestTicket not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	idSource := types.NewPID(0)
	err = idSource.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RequestTicket(fmt.Errorf("Failed to read idSource from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	idTarget := types.NewPID(0)
	err = idTarget.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RequestTicket(fmt.Errorf("Failed to read idTarget from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RequestTicket(nil, packet, callID, idSource, idTarget)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
