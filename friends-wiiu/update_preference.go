// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdatePreference(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UpdatePreference == nil {
		globals.Logger.Warning("FriendsWiiU::UpdatePreference not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	principalPreference, err := parametersStream.ReadStructure(friends_wiiu_types.NewPrincipalPreference())
	if err != nil {
		errorCode = protocol.UpdatePreference(fmt.Errorf("Failed to read principalPreference from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.UpdatePreference(nil, packet, callID, principalPreference.(*friends_wiiu_types.PrincipalPreference))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
