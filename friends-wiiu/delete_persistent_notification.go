// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeletePersistentNotification(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.DeletePersistentNotification == nil {
		globals.Logger.Warning("FriendsWiiU::DeletePersistentNotification not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	persistentNotifications, err := parametersStream.ReadListStructure(friends_wiiu_types.NewPersistentNotification())
	if err != nil {
		_, errorCode = protocol.DeletePersistentNotification(fmt.Errorf("Failed to read persistentNotifications from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DeletePersistentNotification(nil, packet, callID, persistentNotifications.([]*friends_wiiu_types.PersistentNotification))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
