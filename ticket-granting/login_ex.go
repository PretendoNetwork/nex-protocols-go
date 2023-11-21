// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleLoginEx(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.LoginEx == nil {
		globals.Logger.Warning("TicketGranting::LoginEx not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strUserName, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.LoginEx(fmt.Errorf("Failed to read strUserName from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	oExtraData, err := parametersStream.ReadDataHolder()
	if err != nil {
		_, errorCode = protocol.LoginEx(fmt.Errorf("Failed to read oExtraData from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.LoginEx(nil, packet, callID, strUserName, oExtraData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
