// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostSharedData sets the CompletePostSharedData handler function
func (protocol *Protocol) CompletePostSharedData(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreCompletePostSharedDataParam) uint32) {
	protocol.completePostSharedDataHandler = handler
}

func (protocol *Protocol) handleCompletePostSharedData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.completePostSharedDataHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::CompletePostSharedData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_smash_bros_4_types.NewDataStoreCompletePostSharedDataParam())
	if err != nil {
		errorCode = protocol.completePostSharedDataHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.completePostSharedDataHandler(nil, client, callID, param.(*datastore_super_smash_bros_4_types.DataStoreCompletePostSharedDataParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
