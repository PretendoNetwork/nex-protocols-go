package friends_wiiu

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeletePersistentNotification sets the DeletePersistentNotification handler function
func (protocol *FriendsWiiUProtocol) DeletePersistentNotification(handler func(err error, client *nex.Client, callID uint32, notifications []*PersistentNotification)) {
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

	persistentNotifications, err := parametersStream.ReadListStructure(NewPersistentNotification())
	if err != nil {
		go protocol.DeletePersistentNotificationHandler(err, client, callID, nil)
		return
	}

	go protocol.DeletePersistentNotificationHandler(nil, client, callID, persistentNotifications.([]*PersistentNotification))
}
