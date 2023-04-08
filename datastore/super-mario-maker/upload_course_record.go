package datastore_super_mario_maker

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UploadCourseRecord sets the UploadCourseRecord handler function
func (protocol *DataStoreSuperMarioMakerProtocol) UploadCourseRecord(handler func(err error, client *nex.Client, callID uint32, param *DataStoreUploadCourseRecordParam)) {
	protocol.UploadCourseRecordHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleUploadCourseRecord(packet nex.PacketInterface) {
	if protocol.UploadCourseRecordHandler == nil {
		globals.Logger.Warning("DataStoreSMM::UploadCourseRecord not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(NewDataStoreUploadCourseRecordParam())

	if err != nil {
		go protocol.UploadCourseRecordHandler(err, client, callID, nil)
		return
	}

	go protocol.UploadCourseRecordHandler(nil, client, callID, param.(*DataStoreUploadCourseRecordParam))
}
