// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handlePrepareGetObjectV1(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.PrepareGetObjectV1 == nil {
		globals.Logger.Warning("DataStore::PrepareGetObjectV1 not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStorePrepareGetParamV1())
	if err != nil {
		errorCode = protocol.PrepareGetObjectV1(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.PrepareGetObjectV1(nil, packet, callID, param.(*datastore_types.DataStorePrepareGetParamV1))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
