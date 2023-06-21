package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNotificationUrl sets the GetNotificationUrl handler function
func (protocol *DataStoreProtocol) GetNotificationURL(handler func(err error, client *nex.Client, callID uint32, param *DataStoreGetNotificationURLParam)) {
	protocol.GetNotificationURLHandler = handler
}

func (protocol *DataStoreProtocol) HandleGetNotificationURL(packet nex.PacketInterface) {
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

	param, err := parametersStream.ReadStructure(NewDataStoreGetNotificationURLParam())
	if err != nil {
		go protocol.GetNotificationURLHandler(fmt.Errorf("Failed to read dataStoreGetNotificationURLParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetNotificationURLHandler(nil, client, callID, param.(*DataStoreGetNotificationURLParam))
}
