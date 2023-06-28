package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RateObject sets the RateObject handler function
func (protocol *DataStoreProtocol) RateObject(handler func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, param *datastore_types.DataStoreRateObjectParam, fetchRatings bool)) {
	protocol.RateObjectHandler = handler
}

func (protocol *DataStoreProtocol) HandleRateObject(packet nex.PacketInterface) {
	if protocol.RateObjectHandler == nil {
		globals.Logger.Warning("DataStore::RateObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStructure(datastore_types.NewDataStoreRatingTarget())
	if err != nil {
		go protocol.RateObjectHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), client, callID, nil, nil, false)
		return
	}

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreRateObjectParam())
	if err != nil {
		go protocol.RateObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil, nil, false)
		return
	}

	fetchRatings, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.RateObjectHandler(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), client, callID, nil, nil, false)
		return
	}

	go protocol.RateObjectHandler(nil, client, callID, target.(*datastore_types.DataStoreRatingTarget), param.(*datastore_types.DataStoreRateObjectParam), fetchRatings)
}
