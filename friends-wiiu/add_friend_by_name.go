package friends_wiiu

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriendByName sets the AddFriendByName handler function
func (protocol *FriendsWiiUProtocol) AddFriendByName(handler func(err error, client *nex.Client, callID uint32, username string)) {
	protocol.AddFriendByNameHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleAddFriendByName(packet nex.PacketInterface) {
	if protocol.AddFriendByNameHandler == nil {
		globals.Logger.Warning("FriendsWiiU::AddFriendByName not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	username, err := parametersStream.ReadString()
	if err != nil {
		go protocol.AddFriendByNameHandler(fmt.Errorf("Failed to read username from parameters. %s", err.Error()), client, callID, "")
		return
	}

	go protocol.AddFriendByNameHandler(nil, client, callID, username)
}
