// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSpecificMetaV1 sets the GetSpecificMetaV1 handler function
func (protocol *Protocol) GetSpecificMetaV1(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParamV1) uint32) {
	protocol.getSpecificMetaV1Handler = handler
}

func (protocol *Protocol) handleGetSpecificMetaV1(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getSpecificMetaV1Handler == nil {
		globals.Logger.Warning("DataStore::GetSpecificMetaV1 not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetSpecificMetaParamV1())
	if err != nil {
		errorCode = protocol.getSpecificMetaV1Handler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getSpecificMetaV1Handler(nil, packet, callID, param.(*datastore_types.DataStoreGetSpecificMetaParamV1))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
