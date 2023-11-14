// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddBlackList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AddBlackList == nil {
		globals.Logger.Warning("FriendsWiiU::AddBlackList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	blacklistedPrincipal, err := parametersStream.ReadStructure(friends_wiiu_types.NewBlacklistedPrincipal())
	if err != nil {
		errorCode = protocol.AddBlackList(fmt.Errorf("Failed to read blacklistedPrincipal from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.AddBlackList(nil, packet, callID, blacklistedPrincipal.(*friends_wiiu_types.BlacklistedPrincipal))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
