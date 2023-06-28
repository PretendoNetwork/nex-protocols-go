package friends_wiiu

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends/wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeletePersistentNotification sets the DeletePersistentNotification handler function
func (protocol *FriendsWiiUProtocol) DeletePersistentNotification(handler func(err error, client *nex.Client, callID uint32, notifications []*friends_wiiu_types.PersistentNotification)) {
	protocol.DeletePersistentNotificationHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleDeletePersistentNotification(packet nex.PacketInterface) {
	if protocol.DeletePersistentNotificationHandler == nil {
		globals.Logger.Warning("FriendsWiiU::DeletePersistentNotification not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	persistentNotifications, err := parametersStream.ReadListStructure(friends_wiiu_types.NewPersistentNotification())
	if err != nil {
		go protocol.DeletePersistentNotificationHandler(fmt.Errorf("Failed to read persistentNotifications from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.DeletePersistentNotificationHandler(nil, client, callID, persistentNotifications.([]*friends_wiiu_types.PersistentNotification))
}
