// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetMetasMultipleParam(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetMetasMultipleParam == nil {
		globals.Logger.Warning("DataStore::GetMetasMultipleParam not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	params := types.NewList[*datastore_types.DataStoreGetMetaParam]()
	params.Type = datastore_types.NewDataStoreGetMetaParam()
	err = params.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetMetasMultipleParam(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetMetasMultipleParam(nil, packet, callID, params)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
