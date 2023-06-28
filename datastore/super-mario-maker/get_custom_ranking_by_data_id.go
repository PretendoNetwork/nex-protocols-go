package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCustomRankingByDataId sets the GetCustomRankingByDataId handler function
func (protocol *DataStoreSuperMarioMakerProtocol) GetCustomRankingByDataId(handler func(err error, client *nex.Client, callID uint32, dataStoreGetCustomRankingByDataIdParam *datastore_super_mario_maker_types.DataStoreGetCustomRankingByDataIdParam)) {
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

	dataStoreGetCustomRankingByDataIdParam, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreGetCustomRankingByDataIdParam())
	if err != nil {
		go protocol.GetCustomRankingByDataIdHandler(fmt.Errorf("Failed to read dataStoreGetCustomRankingByDataIdParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetCustomRankingByDataIdHandler(nil, client, callID, dataStoreGetCustomRankingByDataIdParam.(*datastore_super_mario_maker_types.DataStoreGetCustomRankingByDataIdParam))
}
