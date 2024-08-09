// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleRateObjects(packet nex.PacketInterface) {
	if protocol.RateObjects == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::RateObjects not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var targets types.List[datastore_types.DataStoreRatingTarget]
	var params types.List[datastore_types.DataStoreRateObjectParam]
	var transactional types.Bool
	var fetchRatings types.Bool

	var err error

	err = targets.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObjects(fmt.Errorf("Failed to read targets from parameters. %s", err.Error()), packet, callID, targets, params, transactional, fetchRatings)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = params.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObjects(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, targets, params, transactional, fetchRatings)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = transactional.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObjects(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), packet, callID, targets, params, transactional, fetchRatings)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = fetchRatings.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObjects(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), packet, callID, targets, params, transactional, fetchRatings)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RateObjects(nil, packet, callID, targets, params, transactional, fetchRatings)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
