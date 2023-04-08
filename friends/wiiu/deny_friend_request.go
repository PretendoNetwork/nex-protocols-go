package friends_wiiu

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DenyFriendRequest sets the DenyFriendRequest handler function
func (protocol *FriendsWiiUProtocol) DenyFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	protocol.DenyFriendRequestHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleDenyFriendRequest(packet nex.PacketInterface) {
	if protocol.DenyFriendRequestHandler == nil {
		globals.Logger.Warning("FriendsWiiU::DenyFriendRequest not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[FriendsWiiU::DenyFriendRequest] Data missing list length")
		go protocol.DenyFriendRequestHandler(err, client, callID, 0)
		return
	}

	id := parametersStream.ReadUInt64LE()

	go protocol.DenyFriendRequestHandler(nil, client, callID, id)
}
