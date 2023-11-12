// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetasWithCourseRecord sets the GetMetasWithCourseRecord handler function
func (protocol *Protocol) GetMetasWithCourseRecord(handler func(err error, packet nex.PacketInterface, callID uint32, params []*datastore_super_mario_maker_types.DataStoreGetCourseRecordParam, metaParam *datastore_types.DataStoreGetMetaParam) uint32) {
	protocol.getMetasWithCourseRecordHandler = handler
}

func (protocol *Protocol) handleGetMetasWithCourseRecord(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getMetasWithCourseRecordHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetMetasWithCourseRecord not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_super_mario_maker_types.NewDataStoreGetCourseRecordParam())
	if err != nil {
		errorCode = protocol.getMetasWithCourseRecordHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	metaParam, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetMetaParam())
	if err != nil {
		errorCode = protocol.getMetasWithCourseRecordHandler(fmt.Errorf("Failed to read metaParam from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getMetasWithCourseRecordHandler(nil, packet, callID, params.([]*datastore_super_mario_maker_types.DataStoreGetCourseRecordParam), metaParam.(*datastore_types.DataStoreGetMetaParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
