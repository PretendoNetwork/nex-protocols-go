// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// MarkFriendRequestsAsReceived sets the MarkFriendRequestsAsReceived handler function
func (protocol *Protocol) MarkFriendRequestsAsReceived(handler func(err error, client *nex.Client, callID uint32, ids []uint64)) {
	protocol.markFriendRequestsAsReceivedHandler = handler
}

func (protocol *Protocol) handleMarkFriendRequestsAsReceived(packet nex.PacketInterface) {
	if protocol.markFriendRequestsAsReceivedHandler == nil {
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
		go protocol.getRequestBlockSettingsHandler(fmt.Errorf("Failed to read ids from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.markFriendRequestsAsReceivedHandler(nil, client, callID, ids)
}
