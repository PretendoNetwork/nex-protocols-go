// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handlePostMetaBinariesWithDataID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.PostMetaBinariesWithDataID == nil {
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
		_, errorCode = protocol.PostMetaBinariesWithDataID(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	params, err := parametersStream.ReadListStructure(datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		_, errorCode = protocol.PostMetaBinariesWithDataID(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	transactional, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.PostMetaBinariesWithDataID(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.PostMetaBinariesWithDataID(nil, packet, callID, dataIDs, params.([]*datastore_types.DataStorePreparePostParam), transactional)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
