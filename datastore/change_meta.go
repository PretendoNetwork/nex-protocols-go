package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ChangeMeta sets the ChangeMeta handler function
func (protocol *DataStoreProtocol) ChangeMeta(handler func(err error, client *nex.Client, callID uint32, dataStoreChangeMetaParam *DataStoreChangeMetaParam)) {
	protocol.ChangeMetaHandler = handler
}

func (protocol *DataStoreProtocol) HandleChangeMeta(packet nex.PacketInterface) {
	if protocol.ChangeMetaHandler == nil {
		globals.Logger.Warning("DataStore::ChangeMeta not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(NewDataStoreChangeMetaParam())
	if err != nil {
		go protocol.ChangeMetaHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.ChangeMetaHandler(nil, client, callID, param.(*DataStoreChangeMetaParam))
}
