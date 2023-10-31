// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchObjectLight sets the SearchObjectLight handler function
func (protocol *Protocol) SearchObjectLight(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreSearchParam) uint32) {
	protocol.searchObjectLightHandler = handler
}

func (protocol *Protocol) handleSearchObjectLight(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.searchObjectLightHandler == nil {
		globals.Logger.Warning("DataStore::SearchObjectLight not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreSearchParam())
	if err != nil {
		errorCode = protocol.searchObjectLightHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.searchObjectLightHandler(nil, packet, callID, param.(*datastore_types.DataStoreSearchParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
