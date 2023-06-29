package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetasWithCourseRecord sets the GetMetasWithCourseRecord handler function
func (protocol *DataStoreSuperMarioMakerProtocol) GetMetasWithCourseRecord(handler func(err error, client *nex.Client, callID uint32, dataStoreGetCourseRecordParams []*datastore_super_mario_maker_types.DataStoreGetCourseRecordParam, dataStoreGetMetaParam *datastore_types.DataStoreGetMetaParam)) {
	protocol.GetMetasWithCourseRecordHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) handleGetMetasWithCourseRecord(packet nex.PacketInterface) {
	if protocol.GetMetasWithCourseRecordHandler == nil {
		globals.Logger.Warning("DataStoreSMM::GetMetasWithCourseRecord not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_super_mario_maker_types.NewDataStoreGetCourseRecordParam())
	if err != nil {
		go protocol.GetMetasWithCourseRecordHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	metaParam, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetMetaParam())
	if err != nil {
		go protocol.GetMetasWithCourseRecordHandler(fmt.Errorf("Failed to read metaParam from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.GetMetasWithCourseRecordHandler(nil, client, callID, params.([]*datastore_super_mario_maker_types.DataStoreGetCourseRecordParam), metaParam.(*datastore_types.DataStoreGetMetaParam))
}
