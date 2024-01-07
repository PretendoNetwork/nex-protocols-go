// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetMetas(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetMetas == nil {
		globals.Logger.Warning("DataStore::GetMetas not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	dataIDs := types.NewList[*types.PrimitiveU64]()
	dataIDs.Type = types.NewPrimitiveU64(0)
	err = dataIDs.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetMetas(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param := datastore_types.NewDataStoreGetMetaParam()
	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetMetas(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetMetas(nil, packet, callID, dataIDs, param)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
