package friends_wiiu

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePreference sets the UpdatePreference handler function
func (protocol *FriendsWiiUProtocol) UpdatePreference(handler func(err error, client *nex.Client, callID uint32, preference *PrincipalPreference)) {
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

	principalPreferenceStructureInterface, err := parametersStream.ReadStructure(NewPrincipalPreference())
	if err != nil {
		go protocol.UpdatePreferenceHandler(err, client, callID, nil)
		return
	}

	principalPreference := principalPreferenceStructureInterface.(*PrincipalPreference)

	go protocol.UpdatePreferenceHandler(nil, client, callID, principalPreference)
}
