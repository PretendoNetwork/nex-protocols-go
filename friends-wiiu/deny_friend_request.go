// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDenyFriendRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.DenyFriendRequest == nil {
		globals.Logger.Warning("FriendsWiiU::DenyFriendRequest not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	id, err := parametersStream.ReadUInt64LE()
	if err != nil {
		_, errorCode = protocol.DenyFriendRequest(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DenyFriendRequest(nil, packet, callID, id)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
