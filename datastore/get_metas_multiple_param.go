// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetasMultipleParam sets the GetMetasMultipleParam handler function
func (protocol *Protocol) GetMetasMultipleParam(handler func(err error, client *nex.Client, callID uint32, params []*datastore_types.DataStoreGetMetaParam) uint32) {
	protocol.getMetasMultipleParamHandler = handler
}

func (protocol *Protocol) handleGetMetasMultipleParam(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getMetasMultipleParamHandler == nil {
		globals.Logger.Warning("DataStore::GetMetasMultipleParam not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreGetMetaParam())
	if err != nil {
		errorCode = protocol.getMetasMultipleParamHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getMetasMultipleParamHandler(nil, client, callID, params.([]*datastore_types.DataStoreGetMetaParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
