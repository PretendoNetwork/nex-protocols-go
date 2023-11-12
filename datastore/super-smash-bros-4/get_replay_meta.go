// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetReplayMeta sets the GetReplayMeta handler function
func (protocol *Protocol) GetReplayMeta(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreGetReplayMetaParam) uint32) {
	protocol.getReplayMetaHandler = handler
}

func (protocol *Protocol) handleGetReplayMeta(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getReplayMetaHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetReplayMeta not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_smash_bros_4_types.NewDataStoreGetReplayMetaParam())
	if err != nil {
		errorCode = protocol.getReplayMetaHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getReplayMetaHandler(nil, packet, callID, param.(*datastore_super_smash_bros_4_types.DataStoreGetReplayMetaParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
