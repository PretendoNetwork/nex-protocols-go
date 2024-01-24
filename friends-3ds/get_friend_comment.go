// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetFriendComment(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetFriendComment == nil {
		globals.Logger.Warning("Friends3DS::GetFriendComment not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	friends := types.NewList[*friends_3ds_types.FriendInfo]()
	friends.Type = friends_3ds_types.NewFriendInfo()
	err = friends.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetFriendComment(fmt.Errorf("Failed to read friends from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetFriendComment(nil, packet, callID, friends)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
