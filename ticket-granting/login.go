// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Login sets the Login handler function
func (protocol *Protocol) Login(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName string) uint32) {
	protocol.loginHandler = handler
}

func (protocol *Protocol) handleLogin(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.loginHandler == nil {
		globals.Logger.Warning("TicketGranting::Login not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strUserName, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.loginHandler(fmt.Errorf("Failed to read strUserName from parameters. %s", err.Error()), packet, callID, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.loginHandler(nil, packet, callID, strUserName)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
