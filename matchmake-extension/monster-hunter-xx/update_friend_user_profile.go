// Package protocol implements the MatchmakeExtensionMonsterHunterXX protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension_monster_hunter_xx_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension/monster-hunter-xx/types"
)

func (protocol *Protocol) handleUpdateFriendUserProfile(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UpdateFriendUserProfile == nil {
		globals.Logger.Warning("MatchmakeExtensionMonsterHunterXX::UpdateFriendUserProfile not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := nex.StreamReadStructure(parametersStream, matchmake_extension_monster_hunter_xx_types.NewFriendUserParam())
	if err != nil {
		_, errorCode = protocol.UpdateFriendUserProfile(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateFriendUserProfile(nil, packet, callID, param)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
