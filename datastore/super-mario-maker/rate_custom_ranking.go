package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RateCustomRanking sets the RateCustomRanking handler function
func (protocol *DataStoreSuperMarioMakerProtocol) RateCustomRanking(handler func(err error, client *nex.Client, callID uint32, dataStoreRateCustomRankingParams []*datastore_super_mario_maker_types.DataStoreRateCustomRankingParam)) {
	protocol.RateCustomRankingHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) handleRateCustomRanking(packet nex.PacketInterface) {
	if protocol.RateCustomRankingHandler == nil {
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
		go protocol.RateCustomRankingHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.RateCustomRankingHandler(nil, client, callID, params.([]*datastore_super_mario_maker_types.DataStoreRateCustomRankingParam))
}
