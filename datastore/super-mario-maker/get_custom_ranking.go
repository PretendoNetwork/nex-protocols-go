// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCustomRanking sets the GetCustomRanking handler function
func (protocol *Protocol) GetCustomRanking(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCustomRankingParam)) {
	protocol.getCustomRankingHandler = handler
}

func (protocol *Protocol) handleGetCustomRanking(packet nex.PacketInterface) {
	if protocol.getCustomRankingHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetCustomRanking not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreGetCustomRankingParam())
	if err != nil {
		go protocol.getCustomRankingHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getCustomRankingHandler(nil, client, callID, param.(*datastore_super_mario_maker_types.DataStoreGetCustomRankingParam))
}
