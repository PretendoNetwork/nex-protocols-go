// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleChangeMetasV1(packet nex.PacketInterface) {
	var err error

	if protocol.ChangeMetasV1 == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::ChangeMetasV1 not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

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
		_, rmcError := protocol.ChangeMetasV1(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	params := types.NewList[*datastore_types.DataStoreChangeMetaParamV1]()
	params.Type = datastore_types.NewDataStoreChangeMetaParamV1()
	err = params.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangeMetasV1(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	transactional := types.NewPrimitiveBool(false)
	err = transactional.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangeMetasV1(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ChangeMetasV1(nil, packet, callID, dataIDs, params, transactional)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
