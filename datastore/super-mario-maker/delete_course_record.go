// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteCourseRecord sets the DeleteCourseRecord handler function
func (protocol *Protocol) DeleteCourseRecord(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCourseRecordParam) uint32) {
	protocol.deleteCourseRecordHandler = handler
}

func (protocol *Protocol) handleDeleteCourseRecord(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deleteCourseRecordHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::DeleteCourseRecord not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreGetCourseRecordParam())
	if err != nil {
		errorCode = protocol.deleteCourseRecordHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.deleteCourseRecordHandler(nil, client, callID, param.(*datastore_super_mario_maker_types.DataStoreGetCourseRecordParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
