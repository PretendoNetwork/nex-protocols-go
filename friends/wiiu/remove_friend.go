package friends_wiiu

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RemoveFriend sets the RemoveFriend handler function
func (protocol *FriendsWiiUProtocol) RemoveFriend(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	protocol.RemoveFriendHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleRemoveFriend(packet nex.PacketInterface) {
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

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsWiiU::RemoveFriend] Data holder not long enough for PID")
		go protocol.RemoveFriendHandler(err, client, callID, 0)
		return
	}

	pid := parametersStream.ReadUInt32LE()

	go protocol.RemoveFriendHandler(nil, client, callID, pid)
}
