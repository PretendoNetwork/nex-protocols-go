// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleBlackListByName(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.BlackListByName == nil {
		globals.Logger.Warning("Friends::BlackListByName not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	strPlayerName, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.BlackListByName(fmt.Errorf("Failed to read strPlayerName from parameters. %s", err.Error()), packet, callID, "", 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiDetails, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.BlackListByName(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), packet, callID, "", 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.BlackListByName(nil, packet, callID, strPlayerName, uiDetails)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
