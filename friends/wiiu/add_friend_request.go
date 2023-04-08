package friends_wiiu

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriendRequest sets the AddFriendRequest handler function
func (protocol *FriendsWiiUProtocol) AddFriendRequest(handler func(err error, client *nex.Client, callID uint32, pid uint32, unknown2 uint8, message string, unknown4 uint8, unknown5 string, gameKey *GameKey, unknown6 *nex.DateTime)) {
	protocol.AddFriendRequestHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleAddFriendRequest(packet nex.PacketInterface) {
	if protocol.AddFriendRequestHandler == nil {
		globals.Logger.Warning("FriendsWiiU::AddFriendRequest not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4+1+1+8 {
		// length check for the following fixed-size data
		// unknown1 + unknown2 + unknown4 + gameKey + unknown6
		err := errors.New("[FriendsWiiU::AddFriendRequest] Data holder not long enough for PID")
		go protocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	pid := parametersStream.ReadUInt32LE()
	unknown2 := parametersStream.ReadUInt8()
	message, err := parametersStream.ReadString()

	if err != nil {
		go protocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	unknown4 := parametersStream.ReadUInt8()
	unknown5, err := parametersStream.ReadString()

	if err != nil {
		go protocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	gameKeyStructureInterface, err := parametersStream.ReadStructure(NewGameKey())
	if err != nil {
		go protocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	gameKey := gameKeyStructureInterface.(*GameKey)

	if err != nil {
		go protocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	unknown6 := nex.NewDateTime(parametersStream.ReadUInt64LE())

	go protocol.AddFriendRequestHandler(nil, client, callID, pid, unknown2, message, unknown4, unknown5, gameKey, unknown6)
}
