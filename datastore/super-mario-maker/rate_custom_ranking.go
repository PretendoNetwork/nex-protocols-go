package datastore_super_mario_maker

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RateCustomRanking sets the RateCustomRanking handler function
func (protocol *DataStoreSuperMarioMakerProtocol) RateCustomRanking(handler func(err error, client *nex.Client, callID uint32, dataStoreRateCustomRankingParams []*DataStoreRateCustomRankingParam)) {
	protocol.RateCustomRankingHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleRateCustomRanking(packet nex.PacketInterface) {
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

	params, err := parametersStream.ReadListStructure(NewDataStoreRateCustomRankingParam())
	if err != nil {
		go protocol.RateCustomRankingHandler(err, client, callID, nil)
		return
	}

	go protocol.RateCustomRankingHandler(nil, client, callID, params.([]*DataStoreRateCustomRankingParam))
}
