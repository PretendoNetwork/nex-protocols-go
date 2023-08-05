// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNotificationURL sets the GetNotificationURL handler function
func (protocol *Protocol) GetNotificationURL(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNotificationURLParam) uint32) {
	protocol.getNotificationURLHandler = handler
}

func (protocol *Protocol) handleGetNotificationURL(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getNotificationURLHandler == nil {
		globals.Logger.Warning("DataStore::GetNotificationURL not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetNotificationURLParam())
	if err != nil {
		errorCode = protocol.getNotificationURLHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getNotificationURLHandler(nil, client, callID, param.(*datastore_types.DataStoreGetNotificationURLParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
