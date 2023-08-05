// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeletePersistentNotification sets the DeletePersistentNotification handler function
func (protocol *Protocol) DeletePersistentNotification(handler func(err error, client *nex.Client, callID uint32, notifications []*friends_wiiu_types.PersistentNotification) uint32) {
	protocol.deletePersistentNotificationHandler = handler
}

func (protocol *Protocol) handleDeletePersistentNotification(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deletePersistentNotificationHandler == nil {
		globals.Logger.Warning("FriendsWiiU::DeletePersistentNotification not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	persistentNotifications, err := parametersStream.ReadListStructure(friends_wiiu_types.NewPersistentNotification())
	if err != nil {
		errorCode = protocol.deletePersistentNotificationHandler(fmt.Errorf("Failed to read persistentNotifications from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.deletePersistentNotificationHandler(nil, client, callID, persistentNotifications.([]*friends_wiiu_types.PersistentNotification))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
