// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePreference sets the UpdatePreference handler function
func (protocol *Protocol) UpdatePreference(handler func(err error, client *nex.Client, callID uint32, preference *friends_wiiu_types.PrincipalPreference)) {
	protocol.updatePreferenceHandler = handler
}

func (protocol *Protocol) handleUpdatePreference(packet nex.PacketInterface) {
	if protocol.updatePreferenceHandler == nil {
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
		go protocol.updatePreferenceHandler(fmt.Errorf("Failed to read principalPreference from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updatePreferenceHandler(nil, client, callID, principalPreference.(*friends_wiiu_types.PrincipalPreference))
}
