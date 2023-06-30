// Package datastore_super_smash_bros_4 implements the Super Smash Bros. 4 DataStore NEX protocol
package datastore_super_smash_bros_4

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetReplayMeta sets the GetReplayMeta handler function
func (protocol *DataStoreSuperSmashBros4Protocol) GetReplayMeta(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreGetReplayMetaParam)) {
	protocol.GetReplayMetaHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) handleGetReplayMeta(packet nex.PacketInterface) {
	if protocol.GetReplayMetaHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::GetReplayMeta not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_smash_bros_4_types.NewDataStoreGetReplayMetaParam())
	if err != nil {
		go protocol.GetReplayMetaHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetReplayMetaHandler(nil, client, callID, param.(*datastore_super_smash_bros_4_types.DataStoreGetReplayMetaParam))
}
