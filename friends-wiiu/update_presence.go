// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdatePresence(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdatePresence == nil {
		globals.Logger.Warning("FriendsWiiU::UpdatePresence not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	nintendoPresenceV2 := friends_wiiu_types.NewNintendoPresenceV2()
	err = nintendoPresenceV2.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdatePresence(fmt.Errorf("Failed to read nintendoPresenceV2 from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdatePresence(nil, packet, callID, nintendoPresenceV2)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
