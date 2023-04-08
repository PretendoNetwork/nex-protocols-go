package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareGetReplay sets the PrepareGetReplay handler function
func (protocol *DataStoreSuperSmashBros4Protocol) PrepareGetReplay(handler func(err error, client *nex.Client, callID uint32, param *DataStorePrepareGetReplayParam)) {
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

	param, err := parametersStream.ReadStructure(NewDataStorePrepareGetReplayParam())
	if err != nil {
		go protocol.PrepareGetReplayHandler(err, client, callID, nil)
		return
	}

	go protocol.PrepareGetReplayHandler(nil, client, callID, param.(*DataStorePrepareGetReplayParam))
}
