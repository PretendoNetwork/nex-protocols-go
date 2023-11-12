// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RateCustomRanking sets the RateCustomRanking handler function
func (protocol *Protocol) RateCustomRanking(handler func(err error, packet nex.PacketInterface, callID uint32, params []*datastore_super_mario_maker_types.DataStoreRateCustomRankingParam) uint32) {
	protocol.rateCustomRankingHandler = handler
}

func (protocol *Protocol) handleRateCustomRanking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.rateCustomRankingHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::RateCustomRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_super_mario_maker_types.NewDataStoreRateCustomRankingParam())
	if err != nil {
		errorCode = protocol.rateCustomRankingHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.rateCustomRankingHandler(nil, packet, callID, params.([]*datastore_super_mario_maker_types.DataStoreRateCustomRankingParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
