package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PreparePostReplay sets the PreparePostReplay handler function
func (protocol *DataStoreSuperSmashBros4Protocol) PreparePostReplay(handler func(err error, client *nex.Client, callID uint32, param *DataStorePreparePostReplayParam)) {
	protocol.PreparePostReplayHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandlePreparePostReplay(packet nex.PacketInterface) {
	if protocol.PreparePostReplayHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::PreparePostReplay not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(NewDataStorePreparePostReplayParam())
	if err != nil {
		go protocol.PreparePostReplayHandler(err, client, callID, nil)
		return
	}

	go protocol.PreparePostReplayHandler(nil, client, callID, param.(*DataStorePreparePostReplayParam))
}
