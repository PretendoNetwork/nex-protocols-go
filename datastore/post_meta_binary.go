package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostMetaBinary sets the PostMetaBinary handler function
func (protocol *DataStoreProtocol) PostMetaBinary(handler func(err error, client *nex.Client, callID uint32, dataStorePreparePostParam *datastore_types.DataStorePreparePostParam)) {
	protocol.PostMetaBinaryHandler = handler
}

func (protocol *DataStoreProtocol) handlePostMetaBinary(packet nex.PacketInterface) {
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

	dataStorePreparePostParam, err := parametersStream.ReadStructure(datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		go protocol.PostMetaBinaryHandler(fmt.Errorf("Failed to read dataStorePreparePostParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.PostMetaBinaryHandler(nil, client, callID, dataStorePreparePostParam.(*datastore_types.DataStorePreparePostParam))
}
