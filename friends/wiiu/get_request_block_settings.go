package friends_wiiu

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRequestBlockSettings sets the GetRequestBlockSettings handler function
func (protocol *FriendsWiiUProtocol) GetRequestBlockSettings(handler func(err error, client *nex.Client, callID uint32, unknowns []uint32)) {
	protocol.GetRequestBlockSettingsHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleGetRequestBlockSettings(packet nex.PacketInterface) {
	if protocol.GetRequestBlockSettingsHandler == nil {
		globals.Logger.Warning("FriendsWiiU::GetRequestBlockSettings not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsWiiU::GetRequestBlockSettings] Data missing list length")
		go protocol.GetRequestBlockSettingsHandler(err, client, callID, make([]uint32, 0))
		return
	}

	pids := parametersStream.ReadListUInt32LE()

	go protocol.GetRequestBlockSettingsHandler(nil, client, callID, pids)
}
