// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddFriendByNameWithDetails(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AddFriendByNameWithDetails == nil {
		globals.Logger.Warning("Friends::AddFriendByNameWithDetails not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	uiPlayer, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.AddFriendByNameWithDetails(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), packet, callID, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiDetails, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.AddFriendByNameWithDetails(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), packet, callID, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.AddFriendByNameWithDetails(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AddFriendByNameWithDetails(nil, packet, callID, uiPlayer, uiDetails, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
