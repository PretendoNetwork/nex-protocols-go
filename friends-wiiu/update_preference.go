package friends_wiiu

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePreference sets the UpdatePreference handler function
func (protocol *FriendsWiiUProtocol) UpdatePreference(handler func(err error, client *nex.Client, callID uint32, preference *friends_wiiu_types.PrincipalPreference)) {
	protocol.UpdatePreferenceHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleUpdatePreference(packet nex.PacketInterface) {
	if protocol.UpdatePreferenceHandler == nil {
		globals.Logger.Warning("FriendsWiiU::UpdatePreference not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	principalPreference, err := parametersStream.ReadStructure(friends_wiiu_types.NewPrincipalPreference())
	if err != nil {
		go protocol.UpdatePreferenceHandler(fmt.Errorf("Failed to read principalPreference from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.UpdatePreferenceHandler(nil, client, callID, principalPreference.(*friends_wiiu_types.PrincipalPreference))
}
