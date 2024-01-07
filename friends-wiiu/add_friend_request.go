// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddFriendRequest(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.AddFriendRequest == nil {
		globals.Logger.Warning("FriendsWiiU::AddFriendRequest not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	pid := types.NewPID(0)
	err = pid.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown2 := types.NewPrimitiveU8(0)
	err = unknown2.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	message := types.NewString("")
	err = message.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read message from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown4 := types.NewPrimitiveU8(0)
	err = unknown4.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read unknown4 from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown5 := types.NewString("")
	err = unknown5.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read unknown5 from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	gameKey := friends_wiiu_types.NewGameKey()
	err = gameKey.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read gameKey from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown6 := types.NewDateTime(0)
	err = unknown6.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read unknown6 from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AddFriendRequest(nil, packet, callID, pid, unknown2, message, unknown4, unknown5, gameKey, unknown6)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
