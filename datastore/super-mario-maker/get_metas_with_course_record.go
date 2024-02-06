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

	if protocol.GetMetasWithCourseRecord == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::GetMetasWithCourseRecord not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

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
		_, rmcError := protocol.GetMetasWithCourseRecord(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	metaParam := datastore_types.NewDataStoreGetMetaParam()
	err = metaParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetMetasWithCourseRecord(fmt.Errorf("Failed to read metaParam from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetMetasWithCourseRecord(nil, packet, callID, params, metaParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
