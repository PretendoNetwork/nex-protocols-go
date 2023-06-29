package datastore_super_smash_bros_4

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostSharedData sets the CompletePostSharedData handler function
func (protocol *DataStoreSuperSmashBros4Protocol) CompletePostSharedData(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStoreCompletePostSharedDataParam)) {
	protocol.CompletePostSharedDataHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) handleCompletePostSharedData(packet nex.PacketInterface) {
	if protocol.CompletePostSharedDataHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::CompletePostSharedData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_smash_bros_4_types.NewDataStoreCompletePostSharedDataParam())
	if err != nil {
		go protocol.CompletePostSharedDataHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.CompletePostSharedDataHandler(nil, client, callID, param.(*datastore_super_smash_bros_4_types.DataStoreCompletePostSharedDataParam))
}
