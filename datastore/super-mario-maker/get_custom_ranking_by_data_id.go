package datastore_super_mario_maker

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCustomRankingByDataId sets the GetCustomRankingByDataId handler function
func (protocol *DataStoreSuperMarioMakerProtocol) GetCustomRankingByDataId(handler func(err error, client *nex.Client, callID uint32, dataStoreGetCustomRankingByDataIdParam *DataStoreGetCustomRankingByDataIdParam)) {
	protocol.GetCustomRankingByDataIdHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleGetCustomRankingByDataId(packet nex.PacketInterface) {
	if protocol.GetCustomRankingByDataIdHandler == nil {
		globals.Logger.Warning("DataStoreSMM::GetCustomRankingByDataId not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStoreGetCustomRankingByDataIdParam, err := parametersStream.ReadStructure(NewDataStoreGetCustomRankingByDataIdParam())

	if err != nil {
		go protocol.GetCustomRankingByDataIdHandler(err, client, callID, nil)
		return
	}

	go protocol.GetCustomRankingByDataIdHandler(nil, client, callID, dataStoreGetCustomRankingByDataIdParam.(*DataStoreGetCustomRankingByDataIdParam))
}
