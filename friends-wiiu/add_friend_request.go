// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddFriendRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AddFriendRequest == nil {
		globals.Logger.Warning("FriendsWiiU::AddFriendRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadPID()
	if err != nil {
		errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown2, err := parametersStream.ReadUInt8()
	if err != nil {
		errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	message, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read message from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown4, err := parametersStream.ReadUInt8()
	if err != nil {
		errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read unknown4 from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown5, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read unknown5 from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	gameKey, err := parametersStream.ReadStructure(friends_wiiu_types.NewGameKey())
	if err != nil {
		errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read gameKey from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown6, err := parametersStream.ReadDateTime()
	if err != nil {
		errorCode = protocol.AddFriendRequest(fmt.Errorf("Failed to read unknown6 from parameters. %s", err.Error()), packet, callID, nil, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.AddFriendRequest(nil, packet, callID, pid, unknown2, message, unknown4, unknown5, gameKey.(*friends_wiiu_types.GameKey), unknown6)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
