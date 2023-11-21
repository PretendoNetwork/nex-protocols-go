// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetFriendMiiList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetFriendMiiList == nil {
		globals.Logger.Warning("Friends3DS::GetFriendMiiList not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	friends, err := nex.StreamReadListStructure(parametersStream, friends_3ds_types.NewFriendInfo())
	if err != nil {
		_, errorCode = protocol.GetFriendMiiList(fmt.Errorf("Failed to read friends from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetFriendMiiList(nil, packet, callID, friends)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
