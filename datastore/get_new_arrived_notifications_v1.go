// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNewArrivedNotificationsV1 sets the GetNewArrivedNotificationsV1 handler function
func (protocol *Protocol) GetNewArrivedNotificationsV1(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam) uint32) {
	protocol.getNewArrivedNotificationsV1Handler = handler
}

func (protocol *Protocol) handleGetNewArrivedNotificationsV1(packet nex.PacketInterface) {
	if protocol.getNewArrivedNotificationsV1Handler == nil {
		globals.Logger.Warning("DataStore::GetNewArrivedNotificationsV1 not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetNewArrivedNotificationsParam())
	if err != nil {
		go protocol.getNewArrivedNotificationsV1Handler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getNewArrivedNotificationsV1Handler(nil, client, callID, param.(*datastore_types.DataStoreGetNewArrivedNotificationsParam))
}
