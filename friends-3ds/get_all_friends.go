// Package protocol implements the Friends 3DS protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAllFriends sets the GetAllFriends handler function
func (protocol *Protocol) GetAllFriends(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.getAllFriendsHandler = handler
}

func (protocol *Protocol) handleGetAllFriends(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getAllFriendsHandler == nil {
		globals.Logger.Warning("Friends3DS::GetAllFriends not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getAllFriendsHandler(nil, client, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
