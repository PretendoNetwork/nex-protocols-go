// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNewArrivedNotificationsV1 sets the GetNewArrivedNotificationsV1 handler function
func (protocol *Protocol) GetNewArrivedNotificationsV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam) uint32) {
	protocol.getNewArrivedNotificationsV1Handler = handler
}

func (protocol *Protocol) handleGetNewArrivedNotificationsV1(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getNewArrivedNotificationsV1Handler == nil {
		globals.Logger.Warning("DataStore::GetNewArrivedNotificationsV1 not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetNewArrivedNotificationsParam())
	if err != nil {
		errorCode = protocol.getNewArrivedNotificationsV1Handler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getNewArrivedNotificationsV1Handler(nil, packet, callID, param.(*datastore_types.DataStoreGetNewArrivedNotificationsParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
