// Package datastore implements the DataStore NEX protocol
package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RateObjects sets the RateObjects handler function
func (protocol *DataStoreProtocol) RateObjects(handler func(err error, client *nex.Client, callID uint32, targets []*datastore_types.DataStoreRatingTarget, params []*datastore_types.DataStoreRateObjectParam, transactional bool, fetchRatings bool)) {
	protocol.RateObjectsHandler = handler
}

func (protocol *DataStoreProtocol) handleRateObjects(packet nex.PacketInterface) {
	if protocol.RateObjectsHandler == nil {
		globals.Logger.Warning("DataStore::RateObjects not implemented")
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
		go protocol.RateObjectsHandler(fmt.Errorf("Failed to read targets from parameters. %s", err.Error()), client, callID, nil, nil, false, false)
		return
	}

	params, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreRateObjectParam())
	if err != nil {
		go protocol.RateObjectsHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil, nil, false, false)
		return
	}

	transactional, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.RateObjectsHandler(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), client, callID, nil, nil, false, false)
		return
	}

	fetchRatings, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.RateObjectsHandler(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), client, callID, nil, nil, false, false)
		return
	}

	go protocol.RateObjectsHandler(nil, client, callID, targets.([]*datastore_types.DataStoreRatingTarget), params.([]*datastore_types.DataStoreRateObjectParam), transactional, fetchRatings)
}
