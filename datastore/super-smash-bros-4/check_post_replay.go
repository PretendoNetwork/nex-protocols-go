package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CheckPostReplay sets the CheckPostReplay handler function
func (protocol *DataStoreSuperSmashBros4Protocol) CheckPostReplay(handler func(err error, client *nex.Client, callID uint32, param *DataStorePreparePostReplayParam)) {
	protocol.CheckPostReplayHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandleCheckPostReplay(packet nex.PacketInterface) {
	if protocol.CheckPostReplayHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::CheckPostReplay not implemented")
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
		go protocol.CheckPostReplayHandler(err, client, callID, nil)
		return
	}

	go protocol.CheckPostReplayHandler(nil, client, callID, param.(*DataStorePreparePostReplayParam))
}
