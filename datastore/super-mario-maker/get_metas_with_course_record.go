package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetasWithCourseRecord sets the GetMetasWithCourseRecord handler function
func (protocol *DataStoreSuperMarioMakerProtocol) GetMetasWithCourseRecord(handler func(err error, client *nex.Client, callID uint32, dataStoreGetCourseRecordParams []*DataStoreGetCourseRecordParam, dataStoreGetMetaParam *datastore.DataStoreGetMetaParam)) {
	protocol.GetMetasWithCourseRecordHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleGetMetasWithCourseRecord(packet nex.PacketInterface) {
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

	params, err := parametersStream.ReadListStructure(NewDataStoreGetCourseRecordParam())
	if err != nil {
		go protocol.GetMetasWithCourseRecordHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	metaParam, err := parametersStream.ReadStructure(datastore.NewDataStoreGetMetaParam())
	if err != nil {
		go protocol.GetMetasWithCourseRecordHandler(fmt.Errorf("Failed to read metaParam from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.GetMetasWithCourseRecordHandler(nil, client, callID, params.([]*DataStoreGetCourseRecordParam), metaParam.(*datastore.DataStoreGetMetaParam))
}
