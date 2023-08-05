// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RateObjects sets the RateObjects handler function
func (protocol *Protocol) RateObjects(handler func(err error, client *nex.Client, callID uint32, targets []*datastore_types.DataStoreRatingTarget, params []*datastore_types.DataStoreRateObjectParam, transactional bool, fetchRatings bool) uint32) {
	protocol.rateObjectsHandler = handler
}

func (protocol *Protocol) handleRateObjects(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.rateObjectsHandler == nil {
		globals.Logger.Warning("DataStore::RateObjects not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	targets, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreRatingTarget())
	if err != nil {
		errorCode = protocol.rateObjectsHandler(fmt.Errorf("Failed to read targets from parameters. %s", err.Error()), client, callID, nil, nil, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	params, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreRateObjectParam())
	if err != nil {
		errorCode = protocol.rateObjectsHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil, nil, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	transactional, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.rateObjectsHandler(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), client, callID, nil, nil, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	fetchRatings, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.rateObjectsHandler(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), client, callID, nil, nil, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.rateObjectsHandler(nil, client, callID, targets.([]*datastore_types.DataStoreRatingTarget), params.([]*datastore_types.DataStoreRateObjectParam), transactional, fetchRatings)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
