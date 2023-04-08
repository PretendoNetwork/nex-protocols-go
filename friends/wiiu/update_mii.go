package friends_wiiu

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMii sets the UpdateMii handler function
func (protocol *FriendsWiiUProtocol) UpdateMii(handler func(err error, client *nex.Client, callID uint32, mii *MiiV2)) {
	protocol.UpdateMiiHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleUpdateMii(packet nex.PacketInterface) {
	if protocol.UpdateMiiHandler == nil {
		globals.Logger.Warning("FriendsWiiU::UpdateMii not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	miiV2StructureInterface, err := parametersStream.ReadStructure(NewMiiV2())
	if err != nil {
		go protocol.UpdateMiiHandler(err, client, callID, nil)
		return
	}

	miiV2 := miiV2StructureInterface.(*MiiV2)

	go protocol.UpdateMiiHandler(nil, client, callID, miiV2)
}
