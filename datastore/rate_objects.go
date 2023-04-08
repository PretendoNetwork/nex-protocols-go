package datastore

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RateObjects sets the RateObjects handler function
func (protocol *DataStoreProtocol) RateObjects(handler func(err error, client *nex.Client, callID uint32, targets []*DataStoreRatingTarget, params []*DataStoreRateObjectParam, transactional bool, fetchRatings bool)) {
	protocol.RateObjectsHandler = handler
}

func (protocol *DataStoreProtocol) HandleRateObjects(packet nex.PacketInterface) {
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

	targets, err := parametersStream.ReadListStructure(NewDataStoreRatingTarget())
	if err != nil {
		go protocol.RateObjectsHandler(err, client, callID, nil, nil, false, false)
		return
	}

	params, err := parametersStream.ReadListStructure(NewDataStoreRateObjectParam())
	if err != nil {
		go protocol.RateObjectsHandler(err, client, callID, nil, nil, false, false)
		return
	}

	transactional := (parametersStream.ReadUInt8() == 1)
	fetchRatings := (parametersStream.ReadUInt8() == 1)

	go protocol.RateObjectsHandler(nil, client, callID, targets.([]*DataStoreRatingTarget), params.([]*DataStoreRateObjectParam), transactional, fetchRatings)
}
