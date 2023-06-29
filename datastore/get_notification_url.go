package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNotificationURL sets the GetNotificationURL handler function
func (protocol *DataStoreProtocol) GetNotificationURL(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNotificationURLParam)) {
	protocol.GetNotificationURLHandler = handler
}

func (protocol *DataStoreProtocol) handleGetNotificationURL(packet nex.PacketInterface) {
	if protocol.GetNotificationURLHandler == nil {
		globals.Logger.Warning("DataStore::GetNotificationURL not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetNotificationURLParam())
	if err != nil {
		go protocol.GetNotificationURLHandler(fmt.Errorf("Failed to read dataStoreGetNotificationURLParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetNotificationURLHandler(nil, client, callID, param.(*datastore_types.DataStoreGetNotificationURLParam))
}
