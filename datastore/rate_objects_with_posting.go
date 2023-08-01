// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RateObjectsWithPosting sets the RateObjectsWithPosting handler function
func (protocol *Protocol) RateObjectsWithPosting(handler func(err error, client *nex.Client, callID uint32, targets []*datastore_types.DataStoreRatingTarget, rateParams []*datastore_types.DataStoreRateObjectParam, postParams []*datastore_types.DataStorePreparePostParam, transactional bool, fetchRatings bool)) {
	protocol.rateObjectsWithPostingHandler = handler
}

func (protocol *Protocol) handleRateObjectsWithPosting(packet nex.PacketInterface) {
	if protocol.rateObjectsWithPostingHandler == nil {
		globals.Logger.Warning("DataStore::RateObjectsWithPosting not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	targets, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreRatingTarget())
	if err != nil {
		go protocol.rateObjectsWithPostingHandler(fmt.Errorf("Failed to read targets from parameters. %s", err.Error()), client, callID, nil, nil, nil, false, false)
		return
	}

	rateParams, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreRateObjectParam())
	if err != nil {
		go protocol.rateObjectsWithPostingHandler(fmt.Errorf("Failed to read rateParams from parameters. %s", err.Error()), client, callID, nil, nil, nil, false, false)
		return
	}

	postParams, err := parametersStream.ReadListStructure(datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		go protocol.rateObjectsWithPostingHandler(fmt.Errorf("Failed to read postParams from parameters. %s", err.Error()), client, callID, nil, nil, nil, false, false)
		return
	}

	transactional, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.rateObjectsWithPostingHandler(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), client, callID, nil, nil, nil, false, false)
		return
	}

	fetchRatings, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.rateObjectsWithPostingHandler(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), client, callID, nil, nil, nil, false, false)
		return
	}

	go protocol.rateObjectsWithPostingHandler(nil, client, callID, targets.([]*datastore_types.DataStoreRatingTarget), rateParams.([]*datastore_types.DataStoreRateObjectParam), postParams.([]*datastore_types.DataStorePreparePostParam), transactional, fetchRatings)
}