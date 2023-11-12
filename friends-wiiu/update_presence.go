// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePresence sets the UpdatePresence handler function
func (protocol *Protocol) UpdatePresence(handler func(err error, packet nex.PacketInterface, callID uint32, presence *friends_wiiu_types.NintendoPresenceV2) uint32) {
	protocol.updatePresenceHandler = handler
}

func (protocol *Protocol) handleUpdatePresence(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updatePresenceHandler == nil {
		globals.Logger.Warning("FriendsWiiU::UpdatePresence not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	nintendoPresenceV2, err := parametersStream.ReadStructure(friends_wiiu_types.NewNintendoPresenceV2())
	if err != nil {
		errorCode = protocol.updatePresenceHandler(fmt.Errorf("Failed to read nintendoPresenceV2 from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updatePresenceHandler(nil, packet, callID, nintendoPresenceV2.(*friends_wiiu_types.NintendoPresenceV2))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
