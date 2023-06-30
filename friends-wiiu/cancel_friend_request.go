// Package friends_wiiu implements the Friends WiiU NEX protocol
package friends_wiiu

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CancelFriendRequest sets the CancelFriendRequest handler function
func (protocol *FriendsWiiUProtocol) CancelFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	protocol.CancelFriendRequestHandler = handler
}

func (protocol *FriendsWiiUProtocol) handleCancelFriendRequest(packet nex.PacketInterface) {
	if protocol.CancelFriendRequestHandler == nil {
		globals.Logger.Warning("FriendsWiiU::CancelFriendRequest not implemented")
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
		go protocol.CancelFriendRequestHandler(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.CancelFriendRequestHandler(nil, client, callID, id)
}
