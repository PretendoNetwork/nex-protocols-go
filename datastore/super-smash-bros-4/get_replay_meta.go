package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetReplayMeta sets the GetReplayMeta handler function
func (protocol *DataStoreSuperSmashBros4Protocol) GetReplayMeta(handler func(err error, client *nex.Client, callID uint32, param *DataStoreGetReplayMetaParam)) {
	protocol.GetReplayMetaHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandleGetReplayMeta(packet nex.PacketInterface) {
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

	param, err := parametersStream.ReadStructure(NewDataStoreGetReplayMetaParam())
	if err != nil {
		go protocol.GetReplayMetaHandler(err, client, callID, nil)
		return
	}

	go protocol.GetReplayMetaHandler(nil, client, callID, param.(*DataStoreGetReplayMetaParam))
}
