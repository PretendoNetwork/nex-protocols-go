// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDenyFriendRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.DenyFriendRequest == nil {
		globals.Logger.Warning("FriendsWiiU::DenyFriendRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[FriendsWiiU::DenyFriendRequest] Data missing list length")
		errorCode = protocol.DenyFriendRequest(err, packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	id, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.DenyFriendRequest(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.DenyFriendRequest(nil, packet, callID, id)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
