// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcceptFriendRequest sets the AcceptFriendRequest handler function
func (protocol *Protocol) AcceptFriendRequest(handler func(err error, packet nex.PacketInterface, callID uint32, id uint64) uint32) {
	protocol.acceptFriendRequestHandler = handler
}

func (protocol *Protocol) handleAcceptFriendRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.acceptFriendRequestHandler == nil {
		globals.Logger.Warning("FriendsWiiU::AcceptFriendRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.acceptFriendRequestHandler(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.acceptFriendRequestHandler(nil, packet, callID, id)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
