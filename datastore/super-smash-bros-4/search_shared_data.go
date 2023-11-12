// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchSharedData sets the SearchSharedData handler function
func (protocol *Protocol) SearchSharedData(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreSearchSharedDataParam) uint32) {
	protocol.searchSharedDataHandler = handler
}

func (protocol *Protocol) handleSearchSharedData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.searchSharedDataHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::SearchSharedData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_smash_bros_4_types.NewDataStoreSearchSharedDataParam())
	if err != nil {
		errorCode = protocol.searchSharedDataHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.searchSharedDataHandler(nil, packet, callID, param.(*datastore_super_smash_bros_4_types.DataStoreSearchSharedDataParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
