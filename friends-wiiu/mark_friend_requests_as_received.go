// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// MarkFriendRequestsAsReceived sets the MarkFriendRequestsAsReceived handler function
func (protocol *Protocol) MarkFriendRequestsAsReceived(handler func(err error, packet nex.PacketInterface, callID uint32, ids []uint64) uint32) {
	protocol.markFriendRequestsAsReceivedHandler = handler
}

func (protocol *Protocol) handleMarkFriendRequestsAsReceived(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.markFriendRequestsAsReceivedHandler == nil {
		globals.Logger.Warning("FriendsWiiU::MarkFriendRequestsAsReceived not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	ids, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.getRequestBlockSettingsHandler(fmt.Errorf("Failed to read ids from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.markFriendRequestsAsReceivedHandler(nil, packet, callID, ids)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
