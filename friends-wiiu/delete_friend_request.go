// Package friends_wiiu implements the Friends WiiU NEX protocol
package friends_wiiu

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteFriendRequest sets the DeleteFriendRequest handler function
func (protocol *FriendsWiiUProtocol) DeleteFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	protocol.DeleteFriendRequestHandler = handler
}

func (protocol *FriendsWiiUProtocol) handleDeleteFriendRequest(packet nex.PacketInterface) {
	if protocol.DeleteFriendRequestHandler == nil {
		globals.Logger.Warning("FriendsWiiU::DeleteFriendRequest not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.DeleteFriendRequestHandler(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.DeleteFriendRequestHandler(nil, client, callID, id)
}
