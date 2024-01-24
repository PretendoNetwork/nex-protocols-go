// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRateObjectWithPosting(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.RateObjectWithPosting == nil {
		globals.Logger.Warning("DataStore::RateObjectWithPosting not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	target := datastore_types.NewDataStoreRatingTarget()
	err = target.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RateObjectWithPosting(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rateParam := datastore_types.NewDataStoreRateObjectParam()
	err = rateParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RateObjectWithPosting(fmt.Errorf("Failed to read rateParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	postParam := datastore_types.NewDataStorePreparePostParam()
	err = postParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RateObjectWithPosting(fmt.Errorf("Failed to read postParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	fetchRatings := types.NewPrimitiveBool(false)
	err = fetchRatings.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RateObjectWithPosting(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RateObjectWithPosting(nil, packet, callID, target, rateParam, postParam, fetchRatings)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
