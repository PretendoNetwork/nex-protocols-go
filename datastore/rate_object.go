// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RateObject sets the RateObject handler function
func (protocol *Protocol) RateObject(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, param *datastore_types.DataStoreRateObjectParam, fetchRatings bool) uint32) {
	protocol.rateObjectHandler = handler
}

func (protocol *Protocol) handleRateObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.rateObjectHandler == nil {
		globals.Logger.Warning("DataStore::RateObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStructure(datastore_types.NewDataStoreRatingTarget())
	if err != nil {
		errorCode = protocol.rateObjectHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreRateObjectParam())
	if err != nil {
		errorCode = protocol.rateObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	fetchRatings, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.rateObjectHandler(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), packet, callID, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.rateObjectHandler(nil, packet, callID, target.(*datastore_types.DataStoreRatingTarget), param.(*datastore_types.DataStoreRateObjectParam), fetchRatings)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
