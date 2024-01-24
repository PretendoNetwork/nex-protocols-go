// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRecommendedCourseSearchObject(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.RecommendedCourseSearchObject == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::RecommendedCourseSearchObject not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	param := datastore_types.NewDataStoreSearchParam()
	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RecommendedCourseSearchObject(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	extraData := types.NewList[*types.String]()
	extraData.Type = types.NewString("")
	err = extraData.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RecommendedCourseSearchObject(fmt.Errorf("Failed to read extraData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RecommendedCourseSearchObject(nil, packet, callID, param, extraData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
