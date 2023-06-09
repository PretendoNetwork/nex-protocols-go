package datastore_super_smash_bros_4

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostReplay sets the CompletePostReplay handler function
func (protocol *DataStoreSuperSmashBros4Protocol) CompletePostReplay(handler func(err error, client *nex.Client, callID uint32, param *DataStoreCompletePostReplayParam)) {
	protocol.CompletePostReplayHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandleCompletePostReplay(packet nex.PacketInterface) {
	if protocol.CompletePostReplayHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::CompletePostReplay not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(NewDataStoreCompletePostReplayParam())
	if err != nil {
		go protocol.CompletePostReplayHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.CompletePostReplayHandler(nil, client, callID, param.(*DataStoreCompletePostReplayParam))
}
