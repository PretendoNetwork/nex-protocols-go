package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchObjectLight sets the SearchObjectLight handler function
func (protocol *DataStoreProtocol) SearchObjectLight(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam)) {
	protocol.SearchObjectLightHandler = handler
}

func (protocol *DataStoreProtocol) handleSearchObjectLight(packet nex.PacketInterface) {
	if protocol.SearchObjectLightHandler == nil {
		globals.Logger.Warning("DataStore::SearchObjectLight not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreSearchParam())
	if err != nil {
		go protocol.SearchObjectLightHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.SearchObjectLightHandler(nil, client, callID, param.(*datastore_types.DataStoreSearchParam))
}
