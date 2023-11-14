// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddFriendByName(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AddFriendByName == nil {
		globals.Logger.Warning("FriendsWiiU::AddFriendByName not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	username, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.AddFriendByName(fmt.Errorf("Failed to read username from parameters. %s", err.Error()), packet, callID, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.AddFriendByName(nil, packet, callID, username)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
