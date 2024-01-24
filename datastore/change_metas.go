// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleChangeMetas(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.ChangeMetas == nil {
		globals.Logger.Warning("DataStore::ChangeMetas not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	dataIDs := types.NewList[*types.PrimitiveU64]()
	dataIDs.Type = types.NewPrimitiveU64(0)
	err = dataIDs.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ChangeMetas(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	params := types.NewList[*datastore_types.DataStoreChangeMetaParam]()
	params.Type = datastore_types.NewDataStoreChangeMetaParam()
	err = params.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ChangeMetas(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	transactional := types.NewPrimitiveBool(false)
	err = transactional.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ChangeMetas(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.ChangeMetas(nil, packet, callID, dataIDs, params, transactional)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
