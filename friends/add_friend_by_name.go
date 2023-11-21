// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddFriendByName(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AddFriendByName == nil {
		globals.Logger.Warning("Friends::AddFriendByName not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strPlayerName, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.AddFriendByName(fmt.Errorf("Failed to read strPlayerName from parameters. %s", err.Error()), packet, callID, "", 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiDetails, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.AddFriendByName(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), packet, callID, "", 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.AddFriendByName(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, "", 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AddFriendByName(nil, packet, callID, strPlayerName, uiDetails, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
