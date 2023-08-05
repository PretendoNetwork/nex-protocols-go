// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriendRequest sets the AddFriendRequest handler function
func (protocol *Protocol) AddFriendRequest(handler func(err error, client *nex.Client, callID uint32, pid uint32, unknown2 uint8, message string, unknown4 uint8, unknown5 string, gameKey *friends_wiiu_types.GameKey, unknown6 *nex.DateTime) uint32) {
	protocol.addFriendRequestHandler = handler
}

func (protocol *Protocol) handleAddFriendRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.addFriendRequestHandler == nil {
		globals.Logger.Warning("FriendsWiiU::AddFriendRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.addFriendRequestHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), client, callID, 0, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}
	unknown2, err := parametersStream.ReadUInt8()
	if err != nil {
		errorCode = protocol.addFriendRequestHandler(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), client, callID, 0, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	message, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.addFriendRequestHandler(fmt.Errorf("Failed to read message from parameters. %s", err.Error()), client, callID, 0, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown4, err := parametersStream.ReadUInt8()
	if err != nil {
		errorCode = protocol.addFriendRequestHandler(fmt.Errorf("Failed to read unknown4 from parameters. %s", err.Error()), client, callID, 0, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown5, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.addFriendRequestHandler(fmt.Errorf("Failed to read unknown5 from parameters. %s", err.Error()), client, callID, 0, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	gameKey, err := parametersStream.ReadStructure(friends_wiiu_types.NewGameKey())
	if err != nil {
		errorCode = protocol.addFriendRequestHandler(fmt.Errorf("Failed to read gameKey from parameters. %s", err.Error()), client, callID, 0, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown6, err := parametersStream.ReadDateTime()
	if err != nil {
		errorCode = protocol.addFriendRequestHandler(fmt.Errorf("Failed to read unknown6 from parameters. %s", err.Error()), client, callID, 0, 0, "", 0, "", nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.addFriendRequestHandler(nil, client, callID, pid, unknown2, message, unknown4, unknown5, gameKey.(*friends_wiiu_types.GameKey), unknown6)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
