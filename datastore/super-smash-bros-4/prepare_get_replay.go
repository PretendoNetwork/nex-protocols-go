// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareGetReplay sets the PrepareGetReplay handler function
func (protocol *Protocol) PrepareGetReplay(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePrepareGetReplayParam) uint32) {
	protocol.prepareGetReplayHandler = handler
}

func (protocol *Protocol) handlePrepareGetReplay(packet nex.PacketInterface) {
	if protocol.prepareGetReplayHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::PrepareGetReplay not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_smash_bros_4_types.NewDataStorePrepareGetReplayParam())
	if err != nil {
		go protocol.prepareGetReplayHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.prepareGetReplayHandler(nil, client, callID, param.(*datastore_super_smash_bros_4_types.DataStorePrepareGetReplayParam))
}
