// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetas sets the GetMetas handler function
func (protocol *Protocol) GetMetas(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64, param *datastore_types.DataStoreGetMetaParam) uint32) {
	protocol.getMetasHandler = handler
}

func (protocol *Protocol) handleGetMetas(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getMetasHandler == nil {
		globals.Logger.Warning("DataStore::GetMetas not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDs, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.getMetasHandler(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), client, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetMetaParam())
	if err != nil {
		errorCode = protocol.getMetasHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getMetasHandler(nil, client, callID, dataIDs, param.(*datastore_types.DataStoreGetMetaParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
