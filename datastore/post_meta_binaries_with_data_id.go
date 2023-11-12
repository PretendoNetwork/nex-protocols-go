// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostMetaBinariesWithDataID sets the PostMetaBinariesWithDataID handler function
func (protocol *Protocol) PostMetaBinariesWithDataID(handler func(err error, packet nex.PacketInterface, callID uint32, dataIDs []uint64, params []*datastore_types.DataStorePreparePostParam, transactional bool) uint32) {
	protocol.postMetaBinariesWithDataIDHandler = handler
}

func (protocol *Protocol) handlePostMetaBinariesWithDataID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.postMetaBinariesWithDataIDHandler == nil {
		globals.Logger.Warning("DataStore::PostMetaBinariesWithDataID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDs, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.postMetaBinariesWithDataIDHandler(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	params, err := parametersStream.ReadListStructure(datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		errorCode = protocol.postMetaBinariesWithDataIDHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	transactional, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.postMetaBinariesWithDataIDHandler(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.postMetaBinariesWithDataIDHandler(nil, packet, callID, dataIDs, params.([]*datastore_types.DataStorePreparePostParam), transactional)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
