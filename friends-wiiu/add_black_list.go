// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddBlackList(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.AddBlackList == nil {
		globals.Logger.Warning("FriendsWiiU::AddBlackList not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	blacklistedPrincipal := friends_wiiu_types.NewBlacklistedPrincipal()
	err = blacklistedPrincipal.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AddBlackList(fmt.Errorf("Failed to read blacklistedPrincipal from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AddBlackList(nil, packet, callID, blacklistedPrincipal)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
