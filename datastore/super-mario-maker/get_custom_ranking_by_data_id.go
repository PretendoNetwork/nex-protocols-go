// Package datastore_super_mario_maker implements the Super Mario Maker DataStore NEX protocol
package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCustomRankingByDataID sets the GetCustomRankingByDataID handler function
func (protocol *DataStoreSuperMarioMakerProtocol) GetCustomRankingByDataID(handler func(err error, client *nex.Client, callID uint32, dataStoreGetCustomRankingByDataIDParam *datastore_super_mario_maker_types.DataStoreGetCustomRankingByDataIDParam)) {
	protocol.GetCustomRankingByDataIDHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) handleGetCustomRankingByDataID(packet nex.PacketInterface) {
	if protocol.GetCustomRankingByDataIDHandler == nil {
		globals.Logger.Warning("DataStoreSMM::GetCustomRankingByDataID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStoreGetCustomRankingByDataIDParam, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreGetCustomRankingByDataIDParam())
	if err != nil {
		go protocol.GetCustomRankingByDataIDHandler(fmt.Errorf("Failed to read dataStoreGetCustomRankingByDataIDParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetCustomRankingByDataIDHandler(nil, client, callID, dataStoreGetCustomRankingByDataIDParam.(*datastore_super_mario_maker_types.DataStoreGetCustomRankingByDataIDParam))
}
