package datastore

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNotificationUrl sets the GetNotificationUrl handler function
func (protocol *DataStoreProtocol) GetNotificationUrl(handler func(err error, client *nex.Client, callID uint32, param *DataStoreGetNotificationUrlParam)) {
	protocol.GetNotificationUrlHandler = handler
}

func (protocol *DataStoreProtocol) HandleGetNotificationUrl(packet nex.PacketInterface) {
	if protocol.GetNotificationUrlHandler == nil {
		globals.Logger.Warning("DataStore::GetNotificationUrl not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(NewDataStoreGetNotificationUrlParam())
	if err != nil {
		go protocol.GetNotificationUrlHandler(err, client, callID, nil)
		return
	}

	go protocol.GetNotificationUrlHandler(nil, client, callID, param.(*DataStoreGetNotificationUrlParam))
}
