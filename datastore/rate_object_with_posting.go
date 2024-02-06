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

	if protocol.RateObjectWithPosting == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::RateObjectWithPosting not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	target := datastore_types.NewDataStoreRatingTarget()
	err = target.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObjectWithPosting(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rateParam := datastore_types.NewDataStoreRateObjectParam()
	err = rateParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObjectWithPosting(fmt.Errorf("Failed to read rateParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	postParam := datastore_types.NewDataStorePreparePostParam()
	err = postParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObjectWithPosting(fmt.Errorf("Failed to read postParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	fetchRatings := types.NewPrimitiveBool(false)
	err = fetchRatings.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObjectWithPosting(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RateObjectWithPosting(nil, packet, callID, target, rateParam, postParam, fetchRatings)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
