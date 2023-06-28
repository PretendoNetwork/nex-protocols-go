package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNewArrivedNotificationsV1 sets the GetNewArrivedNotificationsV1 handler function
func (protocol *DataStoreProtocol) GetNewArrivedNotificationsV1(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam)) {
	protocol.GetNewArrivedNotificationsV1Handler = handler
}

func (protocol *DataStoreProtocol) HandleGetNewArrivedNotificationsV1(packet nex.PacketInterface) {
	if protocol.GetNewArrivedNotificationsV1Handler == nil {
		globals.Logger.Warning("DataStore::GetNewArrivedNotificationsV1 not implemented")
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
		go protocol.GetNewArrivedNotificationsV1Handler(fmt.Errorf("Failed to read dataStoreGetNewArrivedNotificationsParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetNewArrivedNotificationsV1Handler(nil, client, callID, param.(*datastore_types.DataStoreGetNewArrivedNotificationsParam))
}
