package friends_wiiu

import (
	"fmt"

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

	miiV2, err := parametersStream.ReadStructure(NewMiiV2())
	if err != nil {
		go protocol.UpdateMiiHandler(fmt.Errorf("Failed to read miiV2 from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.UpdateMiiHandler(nil, client, callID, miiV2.(*MiiV2))
}
