// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RateObjectWithPosting sets the RateObjectWithPosting handler function
func (protocol *Protocol) RateObjectWithPosting(handler func(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, rateParam *datastore_types.DataStoreRateObjectParam, postParam *datastore_types.DataStorePreparePostParam, fetchRatings bool) uint32) {
	protocol.rateObjectWithPostingHandler = handler
}

func (protocol *Protocol) handleRateObjectWithPosting(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.rateObjectWithPostingHandler == nil {
		globals.Logger.Warning("DataStore::RateObjectWithPosting not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStructure(datastore_types.NewDataStoreRatingTarget())
	if err != nil {
		errorCode = protocol.rateObjectWithPostingHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rateParam, err := parametersStream.ReadStructure(datastore_types.NewDataStoreRateObjectParam())
	if err != nil {
		errorCode = protocol.rateObjectWithPostingHandler(fmt.Errorf("Failed to read rateParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	postParam, err := parametersStream.ReadStructure(datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		errorCode = protocol.rateObjectWithPostingHandler(fmt.Errorf("Failed to read postParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	fetchRatings, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.rateObjectWithPostingHandler(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), packet, callID, nil, nil, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.rateObjectWithPostingHandler(nil, packet, callID, target.(*datastore_types.DataStoreRatingTarget), rateParam.(*datastore_types.DataStoreRateObjectParam), postParam.(*datastore_types.DataStorePreparePostParam), fetchRatings)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
