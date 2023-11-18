// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUploadCourseRecord(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UploadCourseRecord == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::UploadCourseRecord not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := nex.StreamReadStructure(parametersStream, datastore_super_mario_maker_types.NewDataStoreUploadCourseRecordParam())
	if err != nil {
		_, errorCode = protocol.UploadCourseRecord(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UploadCourseRecord(nil, packet, callID, param)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
