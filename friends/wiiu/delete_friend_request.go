package friends_wiiu

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteFriendRequest sets the DeleteFriendRequest handler function
func (protocol *FriendsWiiUProtocol) DeleteFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	protocol.DeleteFriendRequestHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleDeleteFriendRequest(packet nex.PacketInterface) {
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

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[FriendsWiiU::DeleteFriendRequest] Data missing list length")
		go protocol.DeleteFriendRequestHandler(err, client, callID, 0)
		return
	}

	id := parametersStream.ReadUInt64LE()

	go protocol.DeleteFriendRequestHandler(nil, client, callID, id)
}
