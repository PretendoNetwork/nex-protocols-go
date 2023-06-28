package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCourseRecord sets the GetCourseRecord handler function
func (protocol *DataStoreSuperMarioMakerProtocol) GetCourseRecord(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCourseRecordParam)) {
	protocol.GetCourseRecordHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleGetCourseRecord(packet nex.PacketInterface) {
	if protocol.GetCourseRecordHandler == nil {
		globals.Logger.Warning("DataStoreSMM::GetCourseRecord not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreGetCourseRecordParam())
	if err != nil {
		go protocol.GetCourseRecordHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetCourseRecordHandler(nil, client, callID, param.(*datastore_super_mario_maker_types.DataStoreGetCourseRecordParam))
}
