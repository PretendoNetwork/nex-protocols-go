// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostFightingPowerScore sets the PostFightingPowerScore handler function
func (protocol *Protocol) PostFightingPowerScore(handler func(err error, packet nex.PacketInterface, callID uint32, params []*datastore_super_smash_bros_4_types.DataStorePostFightingPowerScoreParam) uint32) {
	protocol.postFightingPowerScoreHandler = handler
}

func (protocol *Protocol) handlePostFightingPowerScore(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.postFightingPowerScoreHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::PostFightingPowerScore not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_super_smash_bros_4_types.NewDataStorePostFightingPowerScoreParam())
	if err != nil {
		errorCode = protocol.postFightingPowerScoreHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.postFightingPowerScoreHandler(nil, packet, callID, params.([]*datastore_super_smash_bros_4_types.DataStorePostFightingPowerScoreParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
