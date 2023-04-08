package friends_wiiu

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// MarkFriendRequestsAsReceived sets the MarkFriendRequestsAsReceived handler function
func (protocol *FriendsWiiUProtocol) MarkFriendRequestsAsReceived(handler func(err error, client *nex.Client, callID uint32, ids []uint64)) {
	protocol.MarkFriendRequestsAsReceivedHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleMarkFriendRequestsAsReceived(packet nex.PacketInterface) {
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

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsWiiU::MarkFriendRequestsAsReceived] Data missing list length")
		go protocol.MarkFriendRequestsAsReceivedHandler(err, client, callID, make([]uint64, 0))
		return
	}

	ids := parametersStream.ReadListUInt64LE()

	go protocol.MarkFriendRequestsAsReceivedHandler(nil, client, callID, ids)
}
