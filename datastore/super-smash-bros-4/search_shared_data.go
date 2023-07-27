// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchSharedData sets the SearchSharedData handler function
func (protocol *Protocol) SearchSharedData(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreSearchSharedDataParam)) {
	protocol.SearchSharedDataHandler = handler
}

func (protocol *Protocol) handleSearchSharedData(packet nex.PacketInterface) {
	if protocol.SearchSharedDataHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::SearchSharedData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_smash_bros_4_types.NewDataStoreSearchSharedDataParam())
	if err != nil {
		go protocol.SearchSharedDataHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.SearchSharedDataHandler(nil, client, callID, param.(*datastore_super_smash_bros_4_types.DataStoreSearchSharedDataParam))
}
