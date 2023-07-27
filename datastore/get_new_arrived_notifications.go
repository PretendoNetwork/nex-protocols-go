// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNewArrivedNotifications sets the GetNewArrivedNotifications handler function
func (protocol *Protocol) GetNewArrivedNotifications(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam)) {
	protocol.getNewArrivedNotificationsHandler = handler
}

func (protocol *Protocol) handleGetNewArrivedNotifications(packet nex.PacketInterface) {
	if protocol.getNewArrivedNotificationsHandler == nil {
		globals.Logger.Warning("DataStore::GetNewArrivedNotifications not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetNewArrivedNotificationsParam())
	if err != nil {
		go protocol.getNewArrivedNotificationsHandler(fmt.Errorf("Failed to read dataStoreGetNewArrivedNotificationsParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getNewArrivedNotificationsHandler(nil, client, callID, param.(*datastore_types.DataStoreGetNewArrivedNotificationsParam))
}
