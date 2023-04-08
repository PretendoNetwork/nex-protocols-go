package friends_wiiu

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RemoveBlackList sets the RemoveBlackList handler function
func (protocol *FriendsWiiUProtocol) RemoveBlackList(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	protocol.RemoveBlackListHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleRemoveBlackList(packet nex.PacketInterface) {
	if protocol.RemoveBlackListHandler == nil {
		globals.Logger.Warning("FriendsWiiU::RemoveBlackList not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsWiiU::RemoveBlackList] Data missing list length")
		go protocol.RemoveBlackListHandler(err, client, callID, 0)
		return
	}

	pid := parametersStream.ReadUInt32LE()

	go protocol.RemoveBlackListHandler(nil, client, callID, pid)
}
