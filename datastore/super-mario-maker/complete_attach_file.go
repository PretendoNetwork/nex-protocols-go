package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompleteAttachFile sets the CompleteAttachFile handler function
func (protocol *DataStoreSuperMarioMakerProtocol) CompleteAttachFile(handler func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParam *datastore_types.DataStoreCompletePostParam)) {
	protocol.CompleteAttachFileHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) handleCompleteAttachFile(packet nex.PacketInterface) {
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

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompletePostParam())
	if err != nil {
		go protocol.CompleteAttachFileHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.CompleteAttachFileHandler(nil, client, callID, param.(*datastore_types.DataStoreCompletePostParam))
}
