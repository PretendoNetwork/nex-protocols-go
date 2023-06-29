package friends_3ds

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAllFriends sets the GetAllFriends handler function
func (protocol *Friends3DSProtocol) GetAllFriends(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetAllFriendsHandler = handler
}

func (protocol *Friends3DSProtocol) handleGetAllFriends(packet nex.PacketInterface) {
	if protocol.GetAllFriendsHandler == nil {
		globals.Logger.Warning("Friends3DS::GetAllFriends not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.GetAllFriendsHandler(nil, client, callID)
}
