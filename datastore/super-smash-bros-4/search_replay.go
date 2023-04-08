package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchReplay sets the SearchReplay handler function
func (protocol *DataStoreSuperSmashBros4Protocol) SearchReplay(handler func(err error, client *nex.Client, callID uint32, param *DataStoreSearchReplayParam)) {
	protocol.SearchReplayHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandleSearchReplay(packet nex.PacketInterface) {
	if protocol.SearchReplayHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::SearchReplay not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(NewDataStoreSearchReplayParam())
	if err != nil {
		go protocol.SearchReplayHandler(err, client, callID, nil)
		return
	}

	go protocol.SearchReplayHandler(nil, client, callID, param.(*DataStoreSearchReplayParam))
}
