// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostFightingPowerScore sets the PostFightingPowerScore handler function
func (protocol *Protocol) PostFightingPowerScore(handler func(err error, client *nex.Client, callID uint32, params []*datastore_super_smash_bros_4_types.DataStorePostFightingPowerScoreParam) uint32) {
	protocol.postFightingPowerScoreHandler = handler
}

func (protocol *Protocol) handlePostFightingPowerScore(packet nex.PacketInterface) {
	if protocol.postFightingPowerScoreHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::PostFightingPowerScore not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_super_smash_bros_4_types.NewDataStorePostFightingPowerScoreParam())
	if err != nil {
		go protocol.postFightingPowerScoreHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.postFightingPowerScoreHandler(nil, client, callID, params.([]*datastore_super_smash_bros_4_types.DataStorePostFightingPowerScoreParam))
}
