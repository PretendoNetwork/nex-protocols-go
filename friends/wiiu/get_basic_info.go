package friends_wiiu

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetBasicInfo sets the GetBasicInfo handler function
func (protocol *FriendsWiiUProtocol) GetBasicInfo(handler func(err error, client *nex.Client, callID uint32, pids []uint32)) {
	protocol.GetBasicInfoHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleGetBasicInfo(packet nex.PacketInterface) {
	if protocol.GetBasicInfoHandler == nil {
		globals.Logger.Warning("FriendsWiiU::GetBasicInfo not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsWiiU::GetBasicInfo] Data missing list length")
		go protocol.GetBasicInfoHandler(err, client, callID, make([]uint32, 0))
		return
	}

	pids := parametersStream.ReadListUInt32LE()

	go protocol.GetBasicInfoHandler(nil, client, callID, pids)
}
