// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostReplay sets the CompletePostReplay handler function
func (protocol *Protocol) CompletePostReplay(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreCompletePostReplayParam) uint32) {
	protocol.completePostReplayHandler = handler
}

func (protocol *Protocol) handleCompletePostReplay(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.completePostReplayHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::CompletePostReplay not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_smash_bros_4_types.NewDataStoreCompletePostReplayParam())
	if err != nil {
		errorCode = protocol.completePostReplayHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.completePostReplayHandler(nil, packet, callID, param.(*datastore_super_smash_bros_4_types.DataStoreCompletePostReplayParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
