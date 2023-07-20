// Package friends implements the Friends QRV protocol
package friends

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcceptFriendship sets the AcceptFriendship handler function
func (protocol *FriendsProtocol) AcceptFriendship(handler func(err error, client *nex.Client, callID uint32, uiPlayer uint32)) {
	protocol.acceptFriendshipHandler = handler
}

func (protocol *FriendsProtocol) handleAcceptFriendship(packet nex.PacketInterface) {
	if protocol.acceptFriendshipHandler == nil {
		globals.Logger.Warning("Friends::AcceptFriendship not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiPlayer, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.acceptFriendshipHandler(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.acceptFriendshipHandler(nil, client, callID, uiPlayer)
}
