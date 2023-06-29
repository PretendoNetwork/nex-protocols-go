package friends_wiiu

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RemoveFriend sets the RemoveFriend handler function
func (protocol *FriendsWiiUProtocol) RemoveFriend(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	protocol.RemoveFriendHandler = handler
}

func (protocol *FriendsWiiUProtocol) handleRemoveFriend(packet nex.PacketInterface) {
	if protocol.RemoveFriendHandler == nil {
		globals.Logger.Warning("FriendsWiiU::RemoveFriend not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.RemoveFriendHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.RemoveFriendHandler(nil, client, callID, pid)
}
