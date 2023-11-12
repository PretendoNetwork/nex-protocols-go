// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateComment sets the UpdateComment handler function
func (protocol *Protocol) UpdateComment(handler func(err error, packet nex.PacketInterface, callID uint32, comment *friends_wiiu_types.Comment) uint32) {
	protocol.updateCommentHandler = handler
}

func (protocol *Protocol) handleUpdateComment(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateCommentHandler == nil {
		globals.Logger.Warning("FriendsWiiU::UpdateComment not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	comment, err := parametersStream.ReadStructure(friends_wiiu_types.NewComment())
	if err != nil {
		errorCode = protocol.updateCommentHandler(fmt.Errorf("Failed to read comment from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateCommentHandler(nil, packet, callID, comment.(*friends_wiiu_types.Comment))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
