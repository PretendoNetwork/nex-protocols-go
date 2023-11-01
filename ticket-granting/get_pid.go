// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPID sets the GetPID handler function
func (protocol *Protocol) GetPID(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName string) uint32) {
	protocol.getPIDHandler = handler
}

func (protocol *Protocol) handleGetPID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPIDHandler == nil {
		globals.Logger.Warning("TicketGranting::GetPID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strUserName, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.getPIDHandler(fmt.Errorf("Failed to read strUserName from parameters. %s", err.Error()), packet, callID, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getPIDHandler(nil, packet, callID, strUserName)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
