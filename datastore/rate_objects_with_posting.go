// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRateObjectsWithPosting(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.RateObjectsWithPosting == nil {
		globals.Logger.Warning("DataStore::RateObjectsWithPosting not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	targets, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreRatingTarget())
	if err != nil {
		_, errorCode = protocol.RateObjectsWithPosting(fmt.Errorf("Failed to read targets from parameters. %s", err.Error()), packet, callID, nil, nil, nil, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rateParams, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreRateObjectParam())
	if err != nil {
		_, errorCode = protocol.RateObjectsWithPosting(fmt.Errorf("Failed to read rateParams from parameters. %s", err.Error()), packet, callID, nil, nil, nil, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	postParams, err := parametersStream.ReadListStructure(datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		_, errorCode = protocol.RateObjectsWithPosting(fmt.Errorf("Failed to read postParams from parameters. %s", err.Error()), packet, callID, nil, nil, nil, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	transactional, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.RateObjectsWithPosting(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), packet, callID, nil, nil, nil, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	fetchRatings, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.RateObjectsWithPosting(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), packet, callID, nil, nil, nil, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RateObjectsWithPosting(nil, packet, callID, targets.([]*datastore_types.DataStoreRatingTarget), rateParams.([]*datastore_types.DataStoreRateObjectParam), postParams.([]*datastore_types.DataStorePreparePostParam), transactional, fetchRatings)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
