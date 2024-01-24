// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdatePreference(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdatePreference == nil {
		globals.Logger.Warning("FriendsWiiU::UpdatePreference not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	principalPreference := friends_wiiu_types.NewPrincipalPreference()
	err = principalPreference.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdatePreference(fmt.Errorf("Failed to read principalPreference from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdatePreference(nil, packet, callID, principalPreference)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
