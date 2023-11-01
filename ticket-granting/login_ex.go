// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// LoginEx sets the LoginEx handler function
func (protocol *Protocol) LoginEx(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName string, oExtraData *nex.DataHolder) uint32) {
	protocol.loginExHandler = handler
}

func (protocol *Protocol) handleLoginEx(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.loginExHandler == nil {
		globals.Logger.Warning("TicketGranting::LoginEx not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strUserName, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.loginExHandler(fmt.Errorf("Failed to read strUserName from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	oExtraData, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.loginExHandler(fmt.Errorf("Failed to read oExtraData from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.loginExHandler(nil, packet, callID, strUserName, oExtraData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
