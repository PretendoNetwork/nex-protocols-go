// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestTicket sets the RequestTicket handler function
func (protocol *Protocol) RequestTicket(handler func(err error, packet nex.PacketInterface, callID uint32, idSource *nex.PID, idTarget *nex.PID) uint32) {
	protocol.requestTicketHandler = handler
}

func (protocol *Protocol) handleRequestTicket(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.requestTicketHandler == nil {
		globals.Logger.Warning("TicketGranting::RequestTicket not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idSource, err := parametersStream.ReadPID()
	if err != nil {
		errorCode = protocol.requestTicketHandler(fmt.Errorf("Failed to read idSource from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	idTarget, err := parametersStream.ReadPID()
	if err != nil {
		errorCode = protocol.requestTicketHandler(fmt.Errorf("Failed to read idTarget from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.requestTicketHandler(nil, packet, callID, idSource, idTarget)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
