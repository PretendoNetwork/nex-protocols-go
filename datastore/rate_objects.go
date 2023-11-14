// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRateObjects(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.RateObjects == nil {
		globals.Logger.Warning("DataStore::RateObjects not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	targets, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreRatingTarget())
	if err != nil {
		errorCode = protocol.RateObjects(fmt.Errorf("Failed to read targets from parameters. %s", err.Error()), packet, callID, nil, nil, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	params, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreRateObjectParam())
	if err != nil {
		errorCode = protocol.RateObjects(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	transactional, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.RateObjects(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), packet, callID, nil, nil, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	fetchRatings, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.RateObjects(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), packet, callID, nil, nil, false, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.RateObjects(nil, packet, callID, targets.([]*datastore_types.DataStoreRatingTarget), params.([]*datastore_types.DataStoreRateObjectParam), transactional, fetchRatings)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
