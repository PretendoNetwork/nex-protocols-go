package datastore_super_smash_bros_4

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareGetReplay sets the PrepareGetReplay handler function
func (protocol *DataStoreSuperSmashBros4Protocol) PrepareGetReplay(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePrepareGetReplayParam)) {
	protocol.PrepareGetReplayHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandlePrepareGetReplay(packet nex.PacketInterface) {
	if protocol.PrepareGetReplayHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::PrepareGetReplay not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_smash_bros_4_types.NewDataStorePrepareGetReplayParam())
	if err != nil {
		go protocol.PrepareGetReplayHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.PrepareGetReplayHandler(nil, client, callID, param.(*datastore_super_smash_bros_4_types.DataStorePrepareGetReplayParam))
}
