package friends_wiiu

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcceptFriendRequest sets the AcceptFriendRequest handler function
func (protocol *FriendsWiiUProtocol) AcceptFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	protocol.AcceptFriendRequestHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleAcceptFriendRequest(packet nex.PacketInterface) {
	if protocol.AcceptFriendRequestHandler == nil {
		globals.Logger.Warning("FriendsWiiU::AcceptFriendRequest not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[FriendsWiiU::AcceptFriendRequest] Data missing list length")
		go protocol.AcceptFriendRequestHandler(err, client, callID, 0)
		return
	}

	id := parametersStream.ReadUInt64LE()

	go protocol.AcceptFriendRequestHandler(nil, client, callID, id)
}
