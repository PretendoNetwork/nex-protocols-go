// Package friends_3ds implements the Friends 3DS NEX protocol
package friends_3ds

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAllFriends sets the GetAllFriends handler function
func (protocol *Friends3DSProtocol) GetAllFriends(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getAllFriendsHandler = handler
}

func (protocol *Friends3DSProtocol) handleGetAllFriends(packet nex.PacketInterface) {
	if protocol.getAllFriendsHandler == nil {
		globals.Logger.Warning("Friends3DS::GetAllFriends not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getAllFriendsHandler(nil, client, callID)
}
