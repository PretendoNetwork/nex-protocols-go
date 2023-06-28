package friends_wiiu

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CheckSettingStatus sets the CheckSettingStatus handler function
func (protocol *FriendsWiiUProtocol) CheckSettingStatus(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.CheckSettingStatusHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleCheckSettingStatus(packet nex.PacketInterface) {
	if protocol.CheckSettingStatusHandler == nil {
		globals.Logger.Warning("FriendsWiiU::CheckSettingStatus not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.CheckSettingStatusHandler(nil, client, callID)
}
