package datastore

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostMetaBinary sets the PostMetaBinary handler function
func (protocol *DataStoreProtocol) PostMetaBinary(handler func(err error, client *nex.Client, callID uint32, dataStorePreparePostParam *DataStorePreparePostParam)) {
	protocol.PostMetaBinaryHandler = handler
}

func (protocol *DataStoreProtocol) HandlePostMetaBinary(packet nex.PacketInterface) {
	if protocol.PostMetaBinaryHandler == nil {
		globals.Logger.Warning("DataStore::PostMetaBinary not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStorePreparePostParam, err := parametersStream.ReadStructure(NewDataStorePreparePostParam())
	if err != nil {
		go protocol.PostMetaBinaryHandler(err, client, callID, nil)
		return
	}

	go protocol.PostMetaBinaryHandler(nil, client, callID, dataStorePreparePostParam.(*DataStorePreparePostParam))
}
