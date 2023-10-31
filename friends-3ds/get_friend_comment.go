// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendComment sets the GetFriendComment handler function
func (protocol *Protocol) GetFriendComment(handler func(err error, packet nex.PacketInterface, callID uint32, friends []*friends_3ds_types.FriendInfo) uint32) {
	protocol.getFriendCommentHandler = handler
}

func (protocol *Protocol) handleGetFriendComment(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getFriendCommentHandler == nil {
		globals.Logger.Warning("Friends3DS::GetFriendComment not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	friends, err := parametersStream.ReadListStructure(friends_3ds_types.NewFriendInfo())
	if err != nil {
		errorCode = protocol.getFriendCommentHandler(fmt.Errorf("Failed to read friends from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getFriendCommentHandler(nil, packet, callID, friends.([]*friends_3ds_types.FriendInfo))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
