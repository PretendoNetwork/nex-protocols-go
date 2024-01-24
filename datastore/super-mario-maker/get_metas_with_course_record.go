// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetMetasWithCourseRecord(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetMetasWithCourseRecord == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetMetasWithCourseRecord not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	params := types.NewList[*datastore_super_mario_maker_types.DataStoreGetCourseRecordParam]()
	params.Type = datastore_super_mario_maker_types.NewDataStoreGetCourseRecordParam()
	err = params.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetMetasWithCourseRecord(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	metaParam := datastore_types.NewDataStoreGetMetaParam()
	err = metaParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetMetasWithCourseRecord(fmt.Errorf("Failed to read metaParam from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetMetasWithCourseRecord(nil, packet, callID, params, metaParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
