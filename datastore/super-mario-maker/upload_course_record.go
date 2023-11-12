// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UploadCourseRecord sets the UploadCourseRecord handler function
func (protocol *Protocol) UploadCourseRecord(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreUploadCourseRecordParam) uint32) {
	protocol.uploadCourseRecordHandler = handler
}

func (protocol *Protocol) handleUploadCourseRecord(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.uploadCourseRecordHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::UploadCourseRecord not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreUploadCourseRecordParam())
	if err != nil {
		errorCode = protocol.uploadCourseRecordHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.uploadCourseRecordHandler(nil, packet, callID, param.(*datastore_super_mario_maker_types.DataStoreUploadCourseRecordParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
