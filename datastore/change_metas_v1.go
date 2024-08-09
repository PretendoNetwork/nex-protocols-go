// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleChangeMetasV1(packet nex.PacketInterface) {
	if protocol.ChangeMetasV1 == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::ChangeMetasV1 not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var dataIDs types.List[types.UInt64]
	var params types.List[datastore_types.DataStoreChangeMetaParamV1]
	var transactional types.Bool

	var err error

	err = dataIDs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangeMetasV1(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), packet, callID, dataIDs, params, transactional)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = params.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangeMetasV1(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, dataIDs, params, transactional)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = transactional.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangeMetasV1(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), packet, callID, dataIDs, params, transactional)
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
