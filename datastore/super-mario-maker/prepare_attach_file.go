package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareAttachFile sets the PrepareAttachFile handler function
func (protocol *DataStoreSuperMarioMakerProtocol) PrepareAttachFile(handler func(err error, client *nex.Client, callID uint32, dataStoreAttachFileParam *datastore_super_mario_maker_types.DataStoreAttachFileParam)) {
	protocol.PrepareAttachFileHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandlePrepareAttachFile(packet nex.PacketInterface) {
	if protocol.PrepareAttachFileHandler == nil {
		globals.Logger.Warning("DataStoreSMM::PrepareAttachFile not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStoreAttachFileParam, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreAttachFileParam())
	if err != nil {
		go protocol.PrepareAttachFileHandler(fmt.Errorf("Failed to read dataStoreAttachFileParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.PrepareAttachFileHandler(nil, client, callID, dataStoreAttachFileParam.(*datastore_super_mario_maker_types.DataStoreAttachFileParam))
}
