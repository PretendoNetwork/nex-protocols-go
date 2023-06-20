package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchObject sets the SearchObject handler function
func (protocol *DataStoreProtocol) SearchObjectLight(handler func(err error, client *nex.Client, callID uint32, param *DataStoreSearchParam)) {
	protocol.SearchObjectLightHandler = handler
}

func (protocol *DataStoreProtocol) HandleSearchObjectLight(packet nex.PacketInterface) {
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

	param, err := parametersStream.ReadStructure(NewDataStoreSearchParam())
	if err != nil {
		go protocol.SearchObjectLightHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.SearchObjectLightHandler(nil, client, callID, param.(*DataStoreSearchParam))
}
