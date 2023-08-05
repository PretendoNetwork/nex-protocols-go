// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchReplay sets the SearchReplay handler function
func (protocol *Protocol) SearchReplay(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreSearchReplayParam) uint32) {
	protocol.searchReplayHandler = handler
}

func (protocol *Protocol) handleSearchReplay(packet nex.PacketInterface) {
	if protocol.searchReplayHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::SearchReplay not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_smash_bros_4_types.NewDataStoreSearchReplayParam())
	if err != nil {
		go protocol.searchReplayHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.searchReplayHandler(nil, client, callID, param.(*datastore_super_smash_bros_4_types.DataStoreSearchReplayParam))
}
