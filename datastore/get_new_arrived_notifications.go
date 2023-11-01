// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNewArrivedNotifications sets the GetNewArrivedNotifications handler function
func (protocol *Protocol) GetNewArrivedNotifications(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam) uint32) {
	protocol.getNewArrivedNotificationsHandler = handler
}

func (protocol *Protocol) handleGetNewArrivedNotifications(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getNewArrivedNotificationsHandler == nil {
		globals.Logger.Warning("DataStore::GetNewArrivedNotifications not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetNewArrivedNotificationsParam())
	if err != nil {
		errorCode = protocol.getNewArrivedNotificationsHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getNewArrivedNotificationsHandler(nil, packet, callID, param.(*datastore_types.DataStoreGetNewArrivedNotificationsParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
