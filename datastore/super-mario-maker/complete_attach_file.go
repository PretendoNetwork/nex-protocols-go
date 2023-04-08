package datastore_super_mario_maker

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompleteAttachFile sets the CompleteAttachFile handler function
func (protocol *DataStoreSuperMarioMakerProtocol) CompleteAttachFile(handler func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParam *datastore.DataStoreCompletePostParam)) {
	protocol.CompleteAttachFileHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleCompleteAttachFile(packet nex.PacketInterface) {
	if protocol.CompleteAttachFileHandler == nil {
		globals.Logger.Warning("DataStoreSMM::CompleteAttachFile not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore.NewDataStoreCompletePostParam())
	if err != nil {
		go protocol.CompleteAttachFileHandler(err, client, callID, nil)
		return
	}

	go protocol.CompleteAttachFileHandler(nil, client, callID, param.(*datastore.DataStoreCompletePostParam))
}
