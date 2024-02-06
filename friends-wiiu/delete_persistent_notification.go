// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeletePersistentNotification(packet nex.PacketInterface) {
	var err error

	if protocol.DeletePersistentNotification == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "FriendsWiiU::DeletePersistentNotification not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	persistentNotifications := types.NewList[*friends_wiiu_types.PersistentNotification]()
	persistentNotifications.Type = friends_wiiu_types.NewPersistentNotification()
	err = persistentNotifications.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeletePersistentNotification(fmt.Errorf("Failed to read persistentNotifications from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeletePersistentNotification(nil, packet, callID, persistentNotifications)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
