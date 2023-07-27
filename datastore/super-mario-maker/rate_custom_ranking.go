// Package protocol implements the Super Mario Maker DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RateCustomRanking sets the RateCustomRanking handler function
func (protocol *Protocol) RateCustomRanking(handler func(err error, client *nex.Client, callID uint32, dataStoreRateCustomRankingParams []*datastore_super_mario_maker_types.DataStoreRateCustomRankingParam)) {
	protocol.rateCustomRankingHandler = handler
}

func (protocol *Protocol) handleRateCustomRanking(packet nex.PacketInterface) {
	if protocol.rateCustomRankingHandler == nil {
		globals.Logger.Warning("DataStoreSMM::RateCustomRanking not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_super_mario_maker_types.NewDataStoreRateCustomRankingParam())
	if err != nil {
		go protocol.rateCustomRankingHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.rateCustomRankingHandler(nil, client, callID, params.([]*datastore_super_mario_maker_types.DataStoreRateCustomRankingParam))
}
