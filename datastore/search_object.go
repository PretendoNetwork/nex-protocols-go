// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchObject sets the SearchObject handler function
func (protocol *Protocol) SearchObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam)) {
	protocol.searchObjectHandler = handler
}

func (protocol *Protocol) handleSearchObject(packet nex.PacketInterface) {
	if protocol.searchObjectHandler == nil {
		globals.Logger.Warning("DataStore::SearchObject not implemented")
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
		go protocol.searchObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.searchObjectHandler(nil, client, callID, param.(*datastore_types.DataStoreSearchParam))
}