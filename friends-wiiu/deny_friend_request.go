// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DenyFriendRequest sets the DenyFriendRequest handler function
func (protocol *Protocol) DenyFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64) uint32) {
	protocol.denyFriendRequestHandler = handler
}

func (protocol *Protocol) handleDenyFriendRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.denyFriendRequestHandler == nil {
		globals.Logger.Warning("FriendsWiiU::DenyFriendRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[FriendsWiiU::DenyFriendRequest] Data missing list length")
		errorCode = protocol.denyFriendRequestHandler(err, client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	id, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.denyFriendRequestHandler(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.denyFriendRequestHandler(nil, client, callID, id)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
