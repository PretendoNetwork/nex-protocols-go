// Package protocol implements the MatchmakeExtensionMonsterHunterXX protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension_monster_hunter_xx_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension/monster-hunter-xx/types"
)

// UpdateFriendUserProfile sets the UpdateFriendUserProfile handler function
func (protocol *Protocol) UpdateFriendUserProfile(handler func(err error, client *nex.Client, callID uint32, param *matchmake_extension_monster_hunter_xx_types.FriendUserParam) uint32) {
	protocol.updateFriendUserProfileHandler = handler
}

func (protocol *Protocol) handleUpdateFriendUserProfile(packet nex.PacketInterface) {
	if protocol.updateFriendUserProfileHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMonsterHunterXX::UpdateFriendUserProfile not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(matchmake_extension_monster_hunter_xx_types.NewFriendUserParam())
	if err != nil {
		go protocol.updateFriendUserProfileHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updateFriendUserProfileHandler(nil, client, callID, param.(*matchmake_extension_monster_hunter_xx_types.FriendUserParam))
}
