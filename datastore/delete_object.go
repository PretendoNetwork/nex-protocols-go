package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteObject sets the DeleteObject handler function
func (protocol *DataStoreProtocol) DeleteObject(handler func(err error, client *nex.Client, callID uint32, param *DataStoreDeleteParam)) {
	protocol.DeleteObjectHandler = handler
}

func (protocol *DataStoreProtocol) HandleDeleteObject(packet nex.PacketInterface) {
	if protocol.DeleteObjectHandler == nil {
		globals.Logger.Warning("DataStore::DeleteObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(NewDataStoreDeleteParam())
	if err != nil {
		go protocol.DeleteObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.DeleteObjectHandler(nil, client, callID, param.(*DataStoreDeleteParam))
}
