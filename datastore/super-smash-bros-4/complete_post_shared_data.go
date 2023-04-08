package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostSharedData sets the CompletePostSharedData handler function
func (protocol *DataStoreSuperSmashBros4Protocol) CompletePostSharedData(handler func(err error, client *nex.Client, callID uint32, param *DataStoreCompletePostSharedDataParam)) {
	protocol.CompletePostSharedDataHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandleCompletePostSharedData(packet nex.PacketInterface) {
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

	param, err := parametersStream.ReadStructure(NewDataStoreCompletePostSharedDataParam())
	if err != nil {
		go protocol.CompletePostSharedDataHandler(err, client, callID, nil)
		return
	}

	go protocol.CompletePostSharedDataHandler(nil, client, callID, param.(*DataStoreCompletePostSharedDataParam))
}
