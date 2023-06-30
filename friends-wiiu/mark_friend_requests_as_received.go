// Package friends_wiiu implements the Friends WiiU NEX protocol
package friends_wiiu

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// MarkFriendRequestsAsReceived sets the MarkFriendRequestsAsReceived handler function
func (protocol *FriendsWiiUProtocol) MarkFriendRequestsAsReceived(handler func(err error, client *nex.Client, callID uint32, ids []uint64)) {
	protocol.MarkFriendRequestsAsReceivedHandler = handler
}

func (protocol *FriendsWiiUProtocol) handleMarkFriendRequestsAsReceived(packet nex.PacketInterface) {
	if protocol.MarkFriendRequestsAsReceivedHandler == nil {
		globals.Logger.Warning("FriendsWiiU::MarkFriendRequestsAsReceived not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	ids, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.GetRequestBlockSettingsHandler(fmt.Errorf("Failed to read ids from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.MarkFriendRequestsAsReceivedHandler(nil, client, callID, ids)
}
