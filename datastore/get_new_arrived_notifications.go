package datastore

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNewArrivedNotifications sets the GetNewArrivedNotifications handler function
func (protocol *DataStoreProtocol) GetNewArrivedNotifications(handler func(err error, client *nex.Client, callID uint32, param *DataStoreGetNewArrivedNotificationsParam)) {
	protocol.GetNewArrivedNotificationsHandler = handler
}

func (protocol *DataStoreProtocol) HandleGetNewArrivedNotifications(packet nex.PacketInterface) {
	if protocol.GetNewArrivedNotificationsHandler == nil {
		globals.Logger.Warning("DataStore::GetNewArrivedNotifications not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(NewDataStoreGetNewArrivedNotificationsParam())
	if err != nil {
		go protocol.GetNewArrivedNotificationsHandler(err, client, callID, nil)
		return
	}

	go protocol.GetNewArrivedNotificationsHandler(nil, client, callID, param.(*DataStoreGetNewArrivedNotificationsParam))
}
