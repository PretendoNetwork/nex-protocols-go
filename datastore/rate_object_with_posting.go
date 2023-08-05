// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RateObjectWithPosting sets the RateObjectWithPosting handler function
func (protocol *Protocol) RateObjectWithPosting(handler func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, rateParam *datastore_types.DataStoreRateObjectParam, postParam *datastore_types.DataStorePreparePostParam, fetchRatings bool) uint32) {
	protocol.rateObjectWithPostingHandler = handler
}

func (protocol *Protocol) handleRateObjectWithPosting(packet nex.PacketInterface) {
	if protocol.rateObjectWithPostingHandler == nil {
		globals.Logger.Warning("DataStore::RateObjectWithPosting not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStructure(datastore_types.NewDataStoreRatingTarget())
	if err != nil {
		go protocol.rateObjectWithPostingHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), client, callID, nil, nil, nil, false)
		return
	}

	rateParam, err := parametersStream.ReadStructure(datastore_types.NewDataStoreRateObjectParam())
	if err != nil {
		go protocol.rateObjectWithPostingHandler(fmt.Errorf("Failed to read rateParam from parameters. %s", err.Error()), client, callID, nil, nil, nil, false)
		return
	}

	postParam, err := parametersStream.ReadStructure(datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		go protocol.rateObjectWithPostingHandler(fmt.Errorf("Failed to read postParam from parameters. %s", err.Error()), client, callID, nil, nil, nil, false)
		return
	}

	fetchRatings, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.rateObjectWithPostingHandler(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), client, callID, nil, nil, nil, false)
		return
	}

	go protocol.rateObjectWithPostingHandler(nil, client, callID, target.(*datastore_types.DataStoreRatingTarget), rateParam.(*datastore_types.DataStoreRateObjectParam), postParam.(*datastore_types.DataStorePreparePostParam), fetchRatings)
}
