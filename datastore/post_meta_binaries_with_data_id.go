// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handlePostMetaBinariesWithDataID(packet nex.PacketInterface) {
	if protocol.PostMetaBinariesWithDataID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::PostMetaBinariesWithDataID not implemented")

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
	params := types.NewList[*datastore_types.DataStorePreparePostParam]()
	params.Type = datastore_types.NewDataStorePreparePostParam()
	transactional := types.NewPrimitiveBool(false)

	var err error

	err = dataIDs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PostMetaBinariesWithDataID(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = params.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PostMetaBinariesWithDataID(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = transactional.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PostMetaBinariesWithDataID(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PostMetaBinariesWithDataID(nil, packet, callID, dataIDs, params, transactional)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
